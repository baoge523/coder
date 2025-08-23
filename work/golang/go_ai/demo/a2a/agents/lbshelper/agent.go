package lbshelper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
	"trpc.group/trpc-go/trpc-go"

	"github.com/cloudwego/eino-ext/callbacks/langfuse"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"trpc.group/trpc-go/trpc-a2a-go/protocol"
	"trpc.group/trpc-go/trpc-a2a-go/taskmanager"
	tgoredis "trpc.group/trpc-go/trpc-database/goredis"
	"trpc.group/trpc-go/trpc-go/log"
)

const (
	defaultStopWord = "TASK_DONE"
)

var ChatModelSystemPrompt = prompt.FromMessages(schema.FString,
	schema.SystemMessage(""))

type state struct {
	Input    *Input `json:"-"`
	Messages []*schema.Message
}

type redisStore struct {
	redisCli redis.UniversalClient
}

func (store *redisStore) Get(ctx context.Context, checkPointID string) ([]byte, bool, error) {
	checkPointKey := fmt.Sprintf("eino_checkpoint:%s", checkPointID)
	val, err := store.redisCli.Get(ctx, checkPointKey).Bytes()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, false, fmt.Errorf("failed to get checkpoint")
	}
	if errors.Is(err, redis.Nil) {
		return nil, false, nil
	}
	return val, true, nil
}

func (store *redisStore) Set(ctx context.Context, checkPointID string, checkPoint []byte) error {
	checkPointKey := fmt.Sprintf("eino_checkpoint:%s", checkPointID)
	_, err := store.redisCli.Set(ctx, checkPointKey, checkPoint, time.Minute*10).Result()
	if err != nil {
		return fmt.Errorf("failed to save checkpoint: %w", err)
	}
	return nil
}

func init() {
	_ = compose.RegisterSerializableType[*state]("lbshelper:state")
	_ = compose.RegisterSerializableType[*Input]("lbshelper:input")
	_ = compose.RegisterSerializableType[*Output]("lbshelper:output")
}

type Input struct {
	UserInput string
}

type Output struct {
}

type Agent struct {
	aMapTools   []tool.BaseTool
	tavilyTools []tool.BaseTool
	askRunnable compose.Runnable[*Input, *Output]
	redisCli    redis.UniversalClient
}

func (a *Agent) GetState(ctx context.Context, checkPointID string) ([]byte, bool, error) {
	checkPointKey := fmt.Sprintf("eino_checkpoint_state:%s", checkPointID)
	val, err := a.redisCli.Get(ctx, checkPointKey).Bytes()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, false, fmt.Errorf("failed to get checkpoint")
	}
	if errors.Is(err, redis.Nil) {
		return nil, false, nil
	}
	return val, true, nil
}

func (a *Agent) SetState(ctx context.Context, checkPointID string, checkPoint []byte) error {
	checkPointKey := fmt.Sprintf("eino_checkpoint_state:%s", checkPointID)
	_, err := a.redisCli.Set(ctx, checkPointKey, checkPoint, time.Minute*10).Result()
	if err != nil {
		return fmt.Errorf("failed to save checkpoint: %w", err)
	}
	return nil
}

func NewAgent() (*Agent, error) {
	a := &Agent{}
	ctx := context.Background()
	cfg := config.GetMainConfig()
	var err error
	a.aMapTools, err = tools.ConnectMCP(ctx, cfg.AMap.ServerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect amap MCP, %w", err)
	}
	a.tavilyTools, err = tools.NewTavily(ctx, &tools.TavilyConfig{APIKey: cfg.Tavily.APIKey})
	if err != nil {
		return nil, fmt.Errorf("failed to connect tavily MCP, %w", err)
	}

	askGraph, err := a.createRunnableGraph(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create graph: %w", err)
	}

	cli, err := tgoredis.New("trpc.redis.lbshelper")
	if err != nil {
		return nil, fmt.Errorf("failed to create redis client: %w", err)
	}

	// 编译 graph，将节点、边、分支转化为面向运行时的结构。由于 graph 中存在环，使用 AnyPredecessor 模式，同时设置运行时最大步数。
	askRunnable, err := askGraph.(*compose.Graph[*Input, *Output]).Compile(ctx,
		compose.WithNodeTriggerMode(compose.AnyPredecessor),
		compose.WithMaxRunSteps(100),
		compose.WithCheckPointStore(&redisStore{redisCli: cli}),
	)
	if err != nil {
		return nil, err
	}
	a.askRunnable = askRunnable
	a.redisCli = cli
	return a, nil
}

func (a *Agent) Process(ctx context.Context, taskID string, initialMsg protocol.Message,
	handle taskmanager.TaskHandler) error {
	part, ok := initialMsg.Parts[0].(protocol.TextPart)
	if !ok {
		return fmt.Errorf("invalid input parts")
	}

	ctx, cancel := context.WithCancel(ctx)
	done := make(chan struct{})
	cb := &callbackHandler{handle: handle, done: done}
	_ = trpc.Go(ctx, time.Minute*10, func(c context.Context) {
		defer cancel()
		select {
		case <-done:
			log.InfoContextf(ctx, "context has canceled")
		case <-ctx.Done():
			log.InfoContextf(ctx, "context has done, err: %v", ctx.Err())
		}
	})

	input := &Input{
		UserInput: part.Text,
	}

	var callbackHandlers []callbacks.Handler
	callbackHandlers = append(callbackHandlers, cb)

	cfg := config.GetMainConfig()
	if cfg.Langfuse.Name != "" {
		cbh, flusher := langfuse.NewLangfuseHandler(&langfuse.Config{
			Host:      cfg.Langfuse.Host,
			PublicKey: cfg.Langfuse.PublicKey,
			SecretKey: cfg.Langfuse.SecretKey,
			Name:      cfg.Langfuse.Name,
			SessionID: taskID,
		})
		defer flusher()
		callbackHandlers = append(callbackHandlers, cbh)
	}

	sr, err := a.askRunnable.Stream(ctx, input,
		compose.WithCheckPointID(taskID),
		compose.WithStateModifier(
			func(ctx context.Context, path compose.NodePath, stateVal any) error {
				stateBytes, exists, err := a.GetState(ctx, taskID)
				if err != nil {
					return fmt.Errorf("failed to get state: %w", err)
				}
				if !exists {
					return nil
				}
				s := stateVal.(*state)
				if err = json.Unmarshal(stateBytes, s); err != nil {
					return err
				}
				s.Input = input
				return nil
			}),
		compose.WithCallbacks(callbackHandlers...))
	if err != nil {
		interruptInfo, ok := compose.ExtractInterruptInfo(err)
		if ok {
			// 保存state
			stateBytes, err := json.Marshal(interruptInfo.State)
			if err != nil {
				return fmt.Errorf("failed to marshal state: %w", err)
			}
			if err = a.SetState(ctx, taskID, stateBytes); err != nil {
				return fmt.Errorf("failed to save state: %w", err)
			}
			if err := handle.UpdateStatus(protocol.TaskStateInputRequired, nil); err != nil {
				log.ErrorContextf(ctx, "failed to update task status, err: %v", err)
			}
			return nil
		}
		if err := handle.UpdateStatus(protocol.TaskStateFailed, nil); err != nil {
			log.ErrorContextf(ctx, "update task status fail, err: %v", err)
		}
		return fmt.Errorf("failed to invoke graph: %w", err)
	}
	defer sr.Close() // remember to close the stream
	for {
		_, err := sr.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				// finish
				break
			}
			if err := handle.UpdateStatus(protocol.TaskStateFailed, nil); err != nil {
				log.ErrorContextf(ctx, "update task status fail, err: %v", err)
			}
			return fmt.Errorf("failed to receive result: %w", err)
		}
	}

	cb.wg.Wait()
	if err = handle.UpdateStatus(protocol.TaskStateCompleted, nil); err != nil {
		log.ErrorContextf(ctx, "update task status fail, err: %v", err)
	}
	return nil
}

func (a *Agent) createRunnableGraph(ctx context.Context, cfg *config.MainConfig) (compose.AnyGraph, error) {
	// 创建一个待编排的 graph，规定整体的输入输出类型，配置全局状态的初始化方法
	graph := compose.NewGraph[*Input, *Output](
		compose.WithGenLocalState(func(ctx context.Context) *state {
			return &state{}
		}),
	)

	startLambda := compose.InvokableLambda(
		func(ctx context.Context, input *Input) (output []*schema.Message, err error) {
			return []*schema.Message{schema.UserMessage(input.UserInput)}, nil
		},
	)
	_ = graph.AddLambdaNode("Lambda:start", startLambda,
		compose.WithNodeName("Lambda:start"),
		compose.WithStatePostHandler(
			func(ctx context.Context, out []*schema.Message, state *state) ([]*schema.Message, error) {
				state.Messages = append(state.Messages, out...)
				return out, nil
			}),
	)

	toolCallingModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  cfg.LLM.APIKey,
		BaseURL: cfg.LLM.URL,
		Model:   cfg.LLM.ChatModel,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create thinking model: %w", err)
	}

	var allTools []tool.BaseTool
	allTools = append(allTools, a.aMapTools...)
	allTools = append(allTools, a.tavilyTools...)

	var toolInfos []*schema.ToolInfo
	// 获取工具信息并绑定到 ChatModel
	for _, v := range allTools {
		var toolInfo *schema.ToolInfo
		toolInfo, err = v.Info(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get tools info, %w", err)
		}
		toolInfos = append(toolInfos, toolInfo)
	}

	err = toolCallingModel.BindTools(toolInfos)
	if err != nil {
		return nil, fmt.Errorf("failed to transfer jina reader tool: %w", err)
	}

	_ = graph.AddChatModelNode("ChatModel:core", toolCallingModel,
		compose.WithNodeName("ChatModel:core"),
		compose.WithStatePreHandler(
			func(ctx context.Context, in []*schema.Message, state *state) ([]*schema.Message, error) {
				systemPrompt, err := ChatModelSystemPrompt.Format(ctx, map[string]any{
					"stop_word": defaultStopWord,
					"meta_info": map[string]interface{}{
						"current_date": time.Now().Format("2006-01-02"),
					},
				})
				if err != nil {
					return nil, fmt.Errorf("failed to format system prompt: %w", err)
				}
				var fullPrompt []*schema.Message
				fullPrompt = append(fullPrompt, systemPrompt...)
				fullPrompt = append(fullPrompt, state.Messages...)
				return fullPrompt, nil
			},
		),
		compose.WithStatePostHandler(
			func(ctx context.Context, out *schema.Message, state *state) (*schema.Message, error) {
				state.Messages = append(state.Messages, out)
				return out, nil
			}),
	)

	toolNode, err := compose.NewToolNode(ctx, &compose.ToolsNodeConfig{
		Tools: allTools,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create tool node: %w", err)
	}
	_ = graph.AddToolsNode("Tool:core", toolNode,
		compose.WithNodeName("Tool:core"),
		compose.WithStatePostHandler(
			func(ctx context.Context, out []*schema.Message, state *state) ([]*schema.Message, error) {
				state.Messages = append(state.Messages, out...)
				return out, nil
			}))

	// wait input
	waitInputLambda := compose.InvokableLambda(
		func(ctx context.Context, input *schema.Message) (output *Input, err error) {
			var userInput *Input
			_ = compose.ProcessState(ctx, func(ctx context.Context, s *state) error {
				userInput = s.Input
				s.Input = nil
				return nil
			})
			if userInput == nil {
				// 中断等待用户输入
				return nil, compose.InterruptAndRerun
			}
			return userInput, nil
		})
	_ = graph.AddLambdaNode("Lambda:wait_input", waitInputLambda,
		compose.WithNodeName("Lambda:wait_input"))

	_ = graph.AddLambdaNode("Lambda:end", compose.InvokableLambda(
		func(ctx context.Context, input *schema.Message) (output *Output, err error) {
			return &Output{}, nil
		}),
		compose.WithNodeName("Lambda:end"))

	_ = graph.AddEdge(compose.START, "Lambda:start")
	_ = graph.AddEdge("Lambda:start", "ChatModel:core")
	_ = graph.AddBranch("ChatModel:core", compose.NewGraphBranch(
		func(ctx context.Context, in *schema.Message) (endNode string, err error) {
			if strings.Contains(in.Content, defaultStopWord) {
				return "Lambda:end", nil
			}
			if len(in.ToolCalls) > 0 {
				return "Tool:core", nil
			}
			return "Lambda:wait_input", nil
		}, map[string]bool{
			"Tool:core":         true,
			"Lambda:wait_input": true,
			"Lambda:end":        true,
		}))

	_ = graph.AddEdge("Lambda:wait_input", "Lambda:start")
	_ = graph.AddEdge("Tool:core", "ChatModel:core")
	_ = graph.AddEdge("Lambda:end", compose.END)

	return graph, nil
}

type callbackHandler struct {
	handle    taskmanager.TaskHandle
	wg        sync.WaitGroup
	done      chan struct{}
	closeOnce sync.Once
}

func (cb *callbackHandler) OnStart(ctx context.Context, info *callbacks.RunInfo,
	input callbacks.CallbackInput) context.Context {
	log.InfoContextf(ctx, "onStart: name=%s, type=%s, compoment=%s", info.Name, info.Type, info.Component)
	if cb.handle == nil {
		return ctx
	}
	cb.wg.Wait()

	switch info.Name {
	case "Tool:core":
		cb.updateWorkingTaskStatus(ctx, "\n> 🔍正在执行网络搜索…\n")
	}

	return ctx
}

func (cb *callbackHandler) OnEnd(ctx context.Context, info *callbacks.RunInfo,
	output callbacks.CallbackOutput) context.Context {
	log.InfoContextf(ctx, "OnEnd: name=%s, type=%s, compoment=%s", info.Name, info.Type, info.Component)
	if cb.handle == nil {
		return ctx
	}

	return ctx
}

func (cb *callbackHandler) OnError(ctx context.Context, info *callbacks.RunInfo, err error) context.Context {
	log.InfoContextf(ctx, "OnError: name=%s, type=%s, compoment=%s", info.Name, info.Type, info.Component)
	return ctx
}

func (cb *callbackHandler) OnStartWithStreamInput(ctx context.Context, info *callbacks.RunInfo,
	input *schema.StreamReader[callbacks.CallbackInput]) context.Context {
	log.InfoContextf(ctx, "OnStartWithStreamInput: name=%s, type=%s, compoment=%s",
		info.Name, info.Type, info.Component)
	defer input.Close()
	return ctx
}

func (cb *callbackHandler) OnEndWithStreamOutput(ctx context.Context, info *callbacks.RunInfo,
	output *schema.StreamReader[callbacks.CallbackOutput]) context.Context {
	log.InfoContextf(ctx, "OnEndWithStreamOutput: name=%s, type=%s, compoment=%s",
		info.Name, info.Type, info.Component)

	switch info.Name {
	case "Tool:core":
		cb.updateWorkingTaskStatus(ctx, "> ✅网络搜索完成\n\n")
	case "ChatModel:core":
		return cb.processChatModelNodeOutput(ctx, info, output)
	}

	return ctx
}

func (cb *callbackHandler) updateWorkingTaskStatus(ctx context.Context, text string) {
	part := protocol.NewTextPart(text)
	err := cb.handle.UpdateStatus(protocol.TaskStateWorking, &protocol.Message{
		Role:  protocol.MessageRoleAgent,
		Parts: []protocol.Part{part},
	})
	if err != nil {
		log.ErrorContextf(ctx, "update status fail, err: %v", err)
		// 任务更新失败把任务取消掉
		cb.closeOnce.Do(func() {
			close(cb.done)
		})
	}
}

func (cb *callbackHandler) processChatModelNodeOutput(ctx context.Context, info *callbacks.RunInfo,
	output *schema.StreamReader[callbacks.CallbackOutput]) context.Context {
	cb.wg.Add(1)
	_ = trpc.Go(ctx, time.Minute*10, func(ctx context.Context) {
		defer cb.wg.Done()
		defer output.Close() // remember to close the stream in defer

		var buffer strings.Builder
		taskDoneMarker := "<" + defaultStopWord + ">"

		for {
			frame, err := output.Recv()
			if errors.Is(err, io.EOF) {
				// finish
				break
			}
			if err != nil {
				log.ErrorContextf(ctx, "processChatModelNodeOutput failed, err: %v", err)
				return
			}
			callbackOutput, ok := frame.(*model.CallbackOutput)
			if !ok {
				log.ErrorContextf(ctx, "invalid message content: %+v", frame)
				return
			}

			// 将当前片段添加到累积内容中
			currentChunk := callbackOutput.Message.Content
			buffer.WriteString(currentChunk)
			fullContent := buffer.String()

			// 检查累积内容中是否有完整的<TASK_DONE>标记
			cleanedContent := strings.ReplaceAll(fullContent, taskDoneMarker, "")

			// 检查是否可以安全发送部分内容
			if len(cleanedContent) > 0 && !mightContainPartialMarker(cleanedContent, taskDoneMarker) {
				cb.updateWorkingTaskStatus(ctx, cleanedContent)
				buffer.Reset()
			}
		}

		// 处理最后可能的内容
		if buffer.Len() > 0 {
			finalContent := buffer.String()
			finalContent = strings.ReplaceAll(finalContent, taskDoneMarker, "")
			if len(finalContent) > 0 {
				cb.updateWorkingTaskStatus(ctx, finalContent)
			}
		}
	})
	return ctx
}

// mightContainPartialMarker 判断content的后缀是否可能包含marker的任意前缀
func mightContainPartialMarker(content, marker string) bool {
	contentLen := len(content)
	markerLen := len(marker)

	// 如果content长度为0，不可能包含marker的任何前缀
	if contentLen == 0 {
		return false
	}

	// 计算需要检查的最大前缀长度
	maxCheckLen := contentLen
	if markerLen < maxCheckLen {
		maxCheckLen = markerLen
	}

	// 检查content的后缀是否是marker某个前缀的后缀
	for i := 1; i <= maxCheckLen; i++ {
		markerPrefix := marker[:i]
		markerPrefixLen := len(markerPrefix)
		contentSuffix := content[max(0, contentLen-markerPrefixLen):]

		if strings.HasSuffix(markerPrefix, contentSuffix) {
			return true
		}
	}

	return false
}
