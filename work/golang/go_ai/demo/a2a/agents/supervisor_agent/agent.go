package supervisor_agent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go_ai/demo/a2a/agents/supervisor_agent/memory"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/cloudwego/eino-ext/callbacks/langfuse"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components"
	einomodel "github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	a2aclient "trpc.group/trpc-go/trpc-a2a-go/client"
	"trpc.group/trpc-go/trpc-a2a-go/protocol"
	"trpc.group/trpc-go/trpc-a2a-go/server"
	"trpc.group/trpc-go/trpc-a2a-go/taskmanager"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

var IntentDetectChatModelSystemPrompt = prompt.FromMessages(schema.GoTemplate,
	schema.SystemMessage(`
你是一个用户助手专家，能够理解用户需求并进行处理

# 注意
- 不要编造回复，对于专业问题你可以使用"send_task"工具，从领域Agent里获取答案回复用户。
- 不要给用户感知到这是一个Multi-Agent系统
- 无论用户以及上下文是什么语言，你需要使用中文输出

# 当前环境信息
当前时间：{{.meta_info.current_date}}

# 可用的Agent列表

{{range .agents}}## Agent 名字: {{.Name}}
Agent 描述: {{.Description}}
Agent 技能:
{{range .Skills}}
- 技能名称: {{.Name}}
  技能描述: {{if .Description}}{{.Description}}{{else}}无描述{{end}}
  示例:
  {{range .Examples}}
    - {{.}}
  {{end}}

{{end}}
{{end}}
`))

type Input struct {
	UserID    string
	UserMsgID string
	Text      string
}

type Output struct {
}

type state struct {
	input        *Input
	agent        *Agent
	mem          memory.Memory
	conversation memory.Conversation
	currentTask  *protocol.Task // 当前任务
}

func NewAgent(opts ...Option) (*Agent, error) {
	opt := defaultOptions()
	for _, o := range opts {
		o(opt)
	}
	a := &Agent{
		isInit: make(chan struct{}),
	}
	ctx := context.Background()
	cfg := config.GetMainConfig()
	var err error
	a.memoryFactory, err = memory.NewRedisMemoryFactory("trpc.redis.supervisor_agent", opt.memoryOptions...)
	if err != nil {
		return nil, fmt.Errorf("failed to create memory factory: %w", err)
	}
	err = a.initAgents(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to init agents: %w", err)
	}
	err = a.createRunnableGraph(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create runnable graph: %w", err)
	}
	return a, nil
}

func (s *state) initMemory(ctx context.Context, userID string) error {
	var err error
	s.mem, err = s.agent.memoryFactory.Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get memory: %w", err)
	}
	s.conversation, err = s.mem.GetCurrentConversation(ctx)
	if err != nil {
		return fmt.Errorf("failed to get current conversation: %w", err)
	}
	s.handleTask(ctx)
	return nil
}

func (s *state) handleTask(ctx context.Context) {
	userState, err := s.mem.GetState(ctx)
	if err != nil {
		log.ErrorContextf(ctx, "failed to get user state: %v", err)
		return
	}
	currentTaskID := userState[memory.StateKeyCurrentTaskID]
	if currentTaskID == "" {
		return
	}
	task, err := s.agent.getTaskInfo(ctx, currentTaskID)
	if err != nil {
		log.ErrorContextf(ctx, "failed to get task info: %v", err)
		return
	}

	// 任务处于等待输入状态，直接将用户输入路由给任务
	if task.Status.State == protocol.TaskStateInputRequired {
		ts, err := time.ParseInLocation(time.RFC3339, task.Status.Timestamp, time.Local)
		if err == nil && time.Now().Sub(ts) <= time.Minute*5 {
			s.currentTask = task // 记录当前任务
		}
		return
	}

	// 新建任务，执行别的任务
	if err := s.mem.SetState(ctx,
		memory.StateKeyCurrentTaskID, "",
	); err != nil {
		log.ErrorContextf(ctx, "failed to update user state: %v", err)
	}
	if !isFinalState(task.Status.State) {
		_, err = s.agent.cancelTask(ctx, currentTaskID)
		if err != nil {
			log.ErrorContextf(ctx, "failed to cancel task: %v", err)
		}
	}
}

// isFinalState checks if a TaskState represents a terminal state.
func isFinalState(state protocol.TaskState) bool {
	return state == protocol.TaskStateCompleted ||
		state == protocol.TaskStateFailed ||
		state == protocol.TaskStateCanceled
}

// Agent Host Agent
type Agent struct {
	runnable            compose.Runnable[*Input, *Output]
	memoryFactory       memory.Factory
	agentClientMap      map[string]*ClientAgent
	tools               []tool.BaseTool
	toolsInfo           []*schema.ToolInfo
	agentInfo           []ClientAgentInfo
	notificationHandler http.Handler
	isInit              chan struct{}
}

type ClientAgentInfo struct {
	Name        string
	Description string
	Skills      []server.AgentSkill
}

func (a *Agent) getTaskInfo(ctx context.Context, idVal string) (*protocol.Task, error) {
	taskID := &model.TaskID{}
	if err := taskID.Decode(idVal); err != nil {
		return nil, fmt.Errorf("failed to decode taskID: %w", err)
	}
	cli, ok := a.agentClientMap[taskID.AgentName]
	if !ok {
		return nil, fmt.Errorf("agent %s not found", taskID.AgentName)
	}
	task, err := cli.a2aClient.GetTasks(ctx, protocol.TaskQueryParams{
		ID: taskID.Encode(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get taskInfo: %w", err)
	}
	return task, nil
}

func (a *Agent) cancelTask(ctx context.Context, idVal string) (*protocol.Task, error) {
	taskID := &model.TaskID{}
	if err := taskID.Decode(idVal); err != nil {
		return nil, fmt.Errorf("failed to decode taskID: %w", err)
	}
	cli, ok := a.agentClientMap[taskID.AgentName]
	if !ok {
		return nil, fmt.Errorf("agent %s not found", taskID.AgentName)
	}
	task, err := cli.a2aClient.CancelTasks(ctx, protocol.TaskIDParams{
		ID: taskID.Encode(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to cancel task: %w", err)
	}
	return task, nil
}

func (a *Agent) updateTask(ctx context.Context, idVal string, msg protocol.Message) (<-chan protocol.StreamingMessageEvent, error) {
	taskID := &model.TaskID{}
	if err := taskID.Decode(idVal); err != nil {
		return nil, fmt.Errorf("failed to decode taskID: %w", err)
	}
	cli, ok := a.agentClientMap[taskID.AgentName]
	if !ok {
		return nil, fmt.Errorf("agent %s not found", taskID.AgentName)
	}
	taskChan, err := cli.a2aClient.StreamMessage(ctx, protocol.SendMessageParams{
		RPCID:   taskID.Encode(),
		Message: msg,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	return taskChan, nil
}

func (a *Agent) initAgentMap(ctx context.Context, cfg *config.MainConfig) error {
	a.agentClientMap = make(map[string]*ClientAgent)
	for _, v := range cfg.HostAgent.Agents {
		clientAgent := &ClientAgent{}
		// 客户端
		a2aClient, err := a2aclient.NewA2AClient(v.ServerURL, a2aclient.WithTimeout(time.Minute*10))
		if err != nil {
			return fmt.Errorf("failed to connect a2a client: %w", err)
		}
		clientAgent.a2aClient = a2aClient

		cardURL := v.CardURL
		if cardURL == "" {
			cardURL = v.ServerURL + defaultAgentCardPath
		}
		// agentCard
		agentCard, err := card.Fetch(cardURL)
		if err != nil {
			return fmt.Errorf("failed to fetch agent card: %w", err)
		}
		clientAgent.agentCard = agentCard
		a.agentClientMap[v.Name] = clientAgent
		a.agentInfo = append(a.agentInfo, ClientAgentInfo{
			Name:        v.Name,
			Description: stringVal(agentCard.Description),
			Skills:      agentCard.Skills,
		})
	}
	return nil
}
func (a *Agent) initAgents(ctx context.Context, cfg *config.MainConfig) error {
	_ = trpc.Go(ctx, time.Minute, func(ctx context.Context) {
		defer func() {
			close(a.isInit)
		}()
		for {
			err := a.initAgentMap(ctx, cfg)
			if err == nil {
				break
			}
			log.ErrorContextf(ctx, "init agent map fail")
			time.Sleep(time.Second)
		}
	})
	sendTaskTool := &SendTask{agent: a}
	clearMemoryTool := &ClearMemory{agent: a}
	a.tools = []tool.BaseTool{sendTaskTool, clearMemoryTool}
	sendTaskToolInfo, err := sendTaskTool.Info(ctx)
	if err != nil {
		return fmt.Errorf("failed to get info of send task tool: %w", err)
	}
	clearMemoryToolInfo, err := clearMemoryTool.Info(ctx)
	if err != nil {
		return fmt.Errorf("failed to get info of clear memory tool: %w", err)
	}
	a.toolsInfo = []*schema.ToolInfo{sendTaskToolInfo, clearMemoryToolInfo}
	return nil
}

func stringVal(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

const (
	defaultAgentCardPath = "/.well-known/agent.json"
	defaultJWKSPath      = "/.well-known/jwks.json"
)

type ClientAgent struct {
	a2aClient *a2aclient.A2AClient
	agentCard *server.AgentCard
}

func (a *Agent) createRunnableGraph(ctx context.Context, cfg *config.MainConfig) error {
	var err error
	// 创建一个待编排的 graph，规定整体的输入输出类型，配置全局状态的初始化方法
	graph := compose.NewGraph[*Input, *Output](
		compose.WithGenLocalState(func(ctx context.Context) *state {
			return &state{agent: a}
		}),
	)

	_ = graph.AddLambdaNode("Lambda:start",
		compose.InvokableLambda(
			func(ctx context.Context, input *Input) (output *state, err error) {

				return &state{}, nil
			}),
		compose.WithNodeName("Lambda:start"),
		compose.WithStatePreHandler(
			func(ctx context.Context, in *Input, state *state) (*Input, error) {
				state.input = in
				if err = state.initMemory(ctx, in.UserID); err != nil {
					return nil, fmt.Errorf("failed init memory: %w", err)
				}
				msg := schema.UserMessage(in.Text)
				msgID := fmt.Sprintf("%s:%s", msg.Role, in.UserMsgID)
				if err = state.conversation.Append(ctx, msgID, msg); err != nil {
					return in, fmt.Errorf("failed to append messages: %w", err)
				}
				if err = state.mem.SetState(ctx,
					memory.StateKeyCurrentUserEventID, in.UserMsgID,
				); err != nil {
					return in, fmt.Errorf("failed to set state: %w", err)
				}
				return in, nil
			}),
		compose.WithStatePostHandler(
			func(ctx context.Context, out *state, state *state) (*state, error) {
				return state, nil
			}),
	)

	_ = graph.AddLambdaNode("Lambda:sendTaskDirectly",
		compose.StreamableLambda(
			func(ctx context.Context, s *state) (*schema.StreamReader[string], error) {
				taskChan, err := a.updateTask(ctx, s.currentTask.ID, protocol.Message{
					Role:  protocol.MessageRoleUser,
					Parts: []protocol.Part{protocol.NewTextPart(s.input.Text)},
				})
				if err != nil {
					return nil, fmt.Errorf("failed to update task: %w", err)
				}
				sr, sw := schema.Pipe[string](1)
				go func() {
					defer sw.Close()

					for event := range taskChan {
						eventBytes, err := json.Marshal(event)
						sw.Send(string(eventBytes), err)
					}
				}()
				return sr, nil
			}),
		compose.WithNodeName("Lambda:sendTaskDirectly"),
	)

	_ = graph.AddLambdaNode("Lambda:sendTaskToEnd",
		compose.InvokableLambda(
			func(ctx context.Context, input string) (*Output, error) {
				return &Output{}, nil
			}),
		compose.WithNodeName("Lambda:sendTaskToEnd"),
	)

	_ = graph.AddLambdaNode("Lambda:toIntent",
		compose.InvokableLambda(
			func(ctx context.Context, s *state) ([]*schema.Message, error) {
				var err error
				messages, err := s.conversation.GetMessages(ctx)
				if err != nil {
					return nil, fmt.Errorf("failed to history messages: %w", err)
				}
				return messages, nil
			}),
		compose.WithNodeName("Lambda:toIntent"),
	)

	intentMode := cfg.LLM.IntentModel
	if intentMode == "" {
		intentMode = cfg.LLM.ChatModel
	}
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  cfg.LLM.APIKey,
		BaseURL: cfg.LLM.URL,
		Model:   intentMode,
	})
	if err != nil {
		return fmt.Errorf("failed to create chat model")
	}

	if err = chatModel.BindTools(a.toolsInfo); err != nil {
		return fmt.Errorf("failed to bind tools: %w", err)
	}

	// 添加意图识别
	_ = graph.AddChatModelNode("ChatModel:intent", chatModel,
		compose.WithNodeName("ChatModel:intent"),
		compose.WithStatePreHandler(
			func(ctx context.Context, in []*schema.Message, state *state) ([]*schema.Message, error) {
				prompts, err := IntentDetectChatModelSystemPrompt.Format(ctx, map[string]any{
					"meta_info": map[string]interface{}{
						"current_date": time.Now().Format("2006-01-02"),
					},
					"agents": a.agentInfo,
				})
				if err != nil {
					return nil, fmt.Errorf("failed to format prompts: %w", err)
				}
				// 拼接用户上下文
				prompts = append(prompts, in...)
				return prompts, nil
			}),
	)

	toolsNode, err := compose.NewToolNode(ctx, &compose.ToolsNodeConfig{
		Tools: a.tools,
	})
	_ = graph.AddToolsNode("Tool:sendTask", toolsNode,
		compose.WithNodeName("Tool:sendTask"),
	)
	_ = graph.AddLambdaNode("Lambda:intentToList", compose.ToList[*schema.Message]())
	_ = graph.AddLambdaNode("Lambda:end",
		compose.InvokableLambda(
			func(ctx context.Context, input []*schema.Message) (*Output, error) {
				return &Output{}, nil
			},
		))

	_ = graph.AddEdge(compose.START, "Lambda:start")
	_ = graph.AddBranch("Lambda:start", compose.NewGraphBranch(
		func(ctx context.Context, in *state) (string, error) {
			if in.currentTask != nil {
				return "Lambda:sendTaskDirectly", nil
			}
			return "Lambda:toIntent", nil
		}, map[string]bool{
			"Lambda:sendTaskDirectly": true,
			"Lambda:toIntent":         true,
		}))

	_ = graph.AddEdge("Lambda:sendTaskDirectly", "Lambda:sendTaskToEnd")
	_ = graph.AddEdge("Lambda:sendTaskToEnd", compose.END)
	_ = graph.AddEdge("Lambda:toIntent", "ChatModel:intent")
	_ = graph.AddBranch("ChatModel:intent", compose.NewGraphBranch(
		func(ctx context.Context, in *schema.Message) (string, error) {
			if len(in.ToolCalls) == 0 {
				return "Lambda:intentToList", nil
			}
			return "Tool:sendTask", nil
		}, map[string]bool{
			"Tool:sendTask":       true,
			"Lambda:intentToList": true,
		}))
	_ = graph.AddEdge("Tool:sendTask", "Lambda:end")
	_ = graph.AddEdge("Lambda:intentToList", "Lambda:end")
	_ = graph.AddEdge("Lambda:end", compose.END)

	// 编译 graph，将节点、边、分支转化为面向运行时的结构。由于 graph 中存在环，使用 AnyPredecessor 模式，同时设置运行时最大步数。
	runnable, err := graph.Compile(ctx,
		compose.WithNodeTriggerMode(compose.AnyPredecessor),
		compose.WithMaxRunSteps(100),
	)
	if err != nil {
		return err
	}
	a.runnable = runnable
	return nil
}

func (a *Agent) Process(ctx context.Context, taskID string, initialMsg protocol.Message,
	handle taskmanager.TaskHandle) error {
	// 等待初始化完成
	<-a.isInit

	part, ok := initialMsg.Parts[0].(protocol.TextPart)
	if !ok {
		return fmt.Errorf("invalid input parts")
	}
	input := &Input{
		UserID:    cast.ToString(initialMsg.Metadata["UserID"]),
		UserMsgID: taskID,
		Text:      part.Text,
	}
	if input.UserID == "" {
		input.UserID = taskID
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

	var callbackHandlers []callbacks.Handler
	callbackHandlers = append(callbackHandlers, cb)

	cfg := config.GetMainConfig()

	// 判断是否配置了langfuse，如果配置了，那么就加入到callbackHandlers中
	if cfg.Langfuse.Name != "" {
		cbh, flusher := langfuse.NewLangfuseHandler(&langfuse.Config{
			Host:      cfg.Langfuse.Host,
			PublicKey: cfg.Langfuse.PublicKey,
			SecretKey: cfg.Langfuse.SecretKey,
			Name:      cfg.Langfuse.Name,
			SessionID: input.UserMsgID,
			UserID:    input.UserID,
		})
		defer flusher()
		callbackHandlers = append(callbackHandlers, cbh)
	}

	// 真正执行AI操作的
	sr, err := a.runnable.Stream(ctx, input, compose.WithCallbacks(callbackHandlers...))
	if err != nil {
		if err := handle.UpdateStatus(protocol.TaskStateFailed, nil); err != nil {
			log.ErrorContextf(ctx, "update task status fail, err: %v", err)
		}
		return fmt.Errorf("failed to invoke graph: %w", err)
	}

	defer sr.Close() // remember to close the stream
	for {
		_, err = sr.Recv()
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

func (a *Agent) handleTaskStatusUpdateEvent(ctx context.Context, event protocol.TaskStatusUpdateEvent) {

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
	if info.Component == components.ComponentOfPrompt {
		return ctx
	}

	cb.wg.Wait()
	var isDone = false

	// 检查用户是否有打断，打断终止任务
	err := compose.ProcessState(ctx, func(ctx context.Context, s *state) error {
		userState, err := s.mem.GetState(ctx)
		if err != nil {
			return fmt.Errorf("failed to get user state: %w", err)
		}
		userMsgID, ok := userState[memory.StateKeyCurrentUserEventID]
		if ok && userMsgID != s.input.UserMsgID {
			log.ErrorContextf(ctx, "user send new message %s", userMsgID)
			isDone = true
		}
		return nil
	})
	if err != nil {
		log.ErrorContextf(ctx, "failed to process state, %v", err)
	}

	if isDone {
		cb.closeOnce.Do(func() {
			close(cb.done)
		})
	}

	return ctx
}

func (cb *callbackHandler) OnEnd(ctx context.Context, info *callbacks.RunInfo,
	output callbacks.CallbackOutput) context.Context {
	log.InfoContextf(ctx, "OnEnd: name=%s, type=%s, compoment=%s", info.Name, info.Type, info.Component)
	return ctx
}

func (cb *callbackHandler) OnError(ctx context.Context, info *callbacks.RunInfo, err error) context.Context {
	log.InfoContextf(ctx, "OnError: name=%s, type=%s, compoment=%s, err=%v",
		info.Name, info.Type, info.Component, err)
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
	case "ChatModel:intent":
		return cb.processIntentNodeOutput(ctx, info, output)
	case "Lambda:sendTaskDirectly":
		return cb.processSendTaskOutput(ctx, info, output)
	case "Tool:sendTask":
		return cb.processSendTaskOutput(ctx, info, output)
	}

	return ctx
}

func (cb *callbackHandler) processIntentNodeOutput(ctx context.Context, info *callbacks.RunInfo,
	output *schema.StreamReader[callbacks.CallbackOutput]) context.Context {
	cb.wg.Add(1)

	// 标识单条消息
	msgID := fmt.Sprintf("%s:%s", schema.Assistant, uuid.New().String())

	_ = trpc.Go(ctx, time.Minute*10, func(ctx context.Context) {
		defer cb.wg.Done()
		defer output.Close() // remember to close the stream in defer

		for {
			frame, err := output.Recv()
			if errors.Is(err, io.EOF) {
				// finish
				break
			}
			if err != nil {
				log.ErrorContextf(ctx, "processIntentNodeOutput failed, err: %v", err)
				return
			}
			callbackOutput, ok := frame.(*einomodel.CallbackOutput)
			if !ok {
				log.ErrorContextf(ctx, "invalid message content: %+v", frame)
				return
			}

			// 只保存不是工具调用的上下文
			if len(callbackOutput.Message.ToolCalls) == 0 {
				err = compose.ProcessState(ctx, func(ctx context.Context, s *state) error {
					return s.conversation.Append(ctx, msgID, callbackOutput.Message)
				})
				if err != nil {
					log.ErrorContextf(ctx, "failed to update conversation, err: %v", err)
					return
				}
			}

			cb.updateWorkingTaskStatus(ctx, callbackOutput.Message.Content)
		}
	})

	return ctx
}

func (cb *callbackHandler) processSendTaskOutput(ctx context.Context, info *callbacks.RunInfo,
	output *schema.StreamReader[callbacks.CallbackOutput]) context.Context {
	cb.wg.Add(1)

	_ = trpc.Go(ctx, time.Minute*10, func(ctx context.Context) {
		defer cb.wg.Done()
		defer output.Close() // remember to close the stream in defer

		for {
			frame, err := output.Recv()
			if errors.Is(err, io.EOF) {
				// finish
				break
			}
			if err != nil {
				log.ErrorContextf(ctx, "processIntentNodeOutput failed, err: %v", err)
				return
			}

			var content string
			switch v := frame.(type) {
			case []*schema.Message:
				if len(v) != 0 {
					content = v[0].Content
				}
			case string:
				content = v
			}
			if content == "" {
				continue
			}
			event := &protocol.TaskStatusUpdateEvent{}
			if err := json.Unmarshal([]byte(content), event); err != nil {
				log.ErrorContextf(ctx, "invalid message content: %+v", content)
				return
			}
			cb.handleTaskStatusUpdateEvent(ctx, event)
		}
	})

	return ctx
}

func (cb *callbackHandler) handleTaskStatusUpdateEvent(ctx context.Context, event *protocol.TaskStatusUpdateEvent) {
	err := compose.ProcessState(ctx, func(ctx context.Context, s *state) error {
		userState, err := s.mem.GetState(ctx)
		if err != nil {
			return fmt.Errorf("failed to get user state: %w", err)
		}
		currentTaskID := userState[memory.StateKeyCurrentTaskID]
		if currentTaskID != event.ID {
			return fmt.Errorf("current task not match")
		}
		currentMsgID := userState[memory.StateKeyCurrentAgentMessageID]
		if currentMsgID == "" {
			currentMsgID = fmt.Sprintf("%s:%s", schema.Assistant, uuid.New().String())
			if err := s.mem.SetState(ctx, memory.StateKeyCurrentAgentMessageID, currentMsgID); err != nil {
				return fmt.Errorf("failed to set current message ID: %w", err)
			}
		}

		if isFinalState(event.Status.State) {
			assistantMsg := schema.AssistantMessage("", nil)
			assistantMsg.ResponseMeta = &schema.ResponseMeta{
				FinishReason: "stop",
			}
			if err = s.conversation.Append(ctx, currentMsgID, assistantMsg); err != nil {
				return fmt.Errorf("failed to update messages: %w", err)
			}

			// 任务已结束
			if err := s.mem.SetState(ctx,
				memory.StateKeyCurrentTaskID, "",
				memory.StateKeyCurrentAgentMessageID, "",
			); err != nil {
				log.ErrorContextf(ctx, "failed to update user state: %v", err)
			}
		}

		if event.Status.State == protocol.TaskStateInputRequired {
			assistantMsg := schema.AssistantMessage("", nil)
			assistantMsg.ResponseMeta = &schema.ResponseMeta{
				FinishReason: string(protocol.TaskStateInputRequired),
			}
			if err = s.conversation.Append(ctx, currentMsgID, assistantMsg); err != nil {
				return fmt.Errorf("failed to update messages: %w", err)
			}
			if err := s.mem.SetState(ctx,
				memory.StateKeyCurrentAgentMessageID, "",
			); err != nil {
				log.ErrorContextf(ctx, "failed to update user state: %v", err)
			}
		}

		if event.Status.Message != nil && len(event.Status.Message.Parts) != 0 {
			part, ok := event.Status.Message.Parts[0].(protocol.TextPart)
			if !ok {
				return fmt.Errorf("unsupported message part")
			}
			if err = s.conversation.Append(ctx, currentMsgID, schema.AssistantMessage(part.Text, nil)); err != nil {
				return fmt.Errorf("failed to update messages: %w", err)
			}
			return nil
		}

		return nil
	})
	if err != nil {
		log.ErrorContextf(ctx, "failed to process state: %v", err)
	}

	if event.Status.Message != nil && len(event.Status.Message.Parts) != 0 {
		if part, ok := event.Status.Message.Parts[0].(protocol.TextPart); ok {
			cb.updateWorkingTaskStatus(ctx, part.Text)
		}
	}
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
