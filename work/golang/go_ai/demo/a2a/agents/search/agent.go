package search

import (
	"a2a-samples/config"
	"a2a-samples/pkg/tools"
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/cloudwego/eino-ext/callbacks/langfuse"
	"github.com/cloudwego/eino-ext/components/model/deepseek"
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"trpc.group/trpc-go/trpc-a2a-go/protocol"
	"trpc.group/trpc-go/trpc-a2a-go/taskmanager"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

const (
	defaultMaxSearchWords = 3
	defaultMaxResult      = 1
)

var AskThinkingModelSystemPrompt = prompt.FromMessages(schema.FString,
	schema.UserMessage(`
你是一个联网信息搜索专家，你需要根据用户的问题，通过联网搜索来搜集相关信息，然后根据这些信息来回答用户的问题。
# 当前环境信息
{meta_info}

# 当前已知资料
{reference}

# 任务
- 判断「当前已知资料」是否已经足够回答用户的问题
- 如果「当前已知资料」已经足够回答用户的问题，返回“无需检索”，不要输出任何其他多余的内容
- 如果判断「当前已知资料」还不足以回答用户的问题，思考还需要搜索什么信息，输出对应的关键词，请保证每个关键词的精简和独立性
- 输出的每个关键词都应该要具体到可以用于独立检索，要包括完整的主语和宾语，避免歧义和使用代词，关键词之间不能有指代关系
- 可以输出1 ~ {max_search_words}个关键词，当暂时无法提出足够准确的关键词时，请适当地减少关键词的数量
- 请略过关键词调优的过程，思考过程里只能有对问题以及已知资料的分析
- 输出多个关键词时，关键词之间用 ; 分割，不要输出其他任何多余的内容
- 你只能输出关键词或者"无需检索"，不能输出其他内容

# 用户问题
{question}

# 你的回答
`))

var AskSummaryModelSystemPrompt = prompt.FromMessages(schema.FString,
	schema.UserMessage(`
# 联网搜索资料
{reference}

# 当前环境信息
{meta_info}

# 任务
- 优先参考「联网参考资料」中的信息进行回复。
- 回复请使用清晰、结构化（序号/分段等）的语言，确保用户轻松理解和使用。
- 如果回复内容中参考了「联网」中的信息，在请务必在正文的段落中引用对应的参考编号，例如[3][5]
- 回答的最后需要列出已参考的所有资料信息。格式如下：[参考编号] 资料名称
示例：
[1] 火山引擎
[3] 火山方舟大模型服务平台

# 任务执行
遵循任务要求来回答「用户问题」，给出有帮助的回答。

# 用户问题
{question}

# 你的回答
`))

type askState struct {
	UserInput string
	Reference referenceArray // 搜索信息
}

type askOutput struct {
	Reference referenceArray
	Content   string
}

type reference struct {
	Keyword string
	Title   string
	URL     string
	Content string
}

type referenceArray []reference

func (arr referenceArray) String() string {
	buf := strings.Builder{}
	for index, ref := range arr {
		buf.WriteString(fmt.Sprintf("----------\n[参考资料 %d 开始]\n", index+1))
		buf.WriteString("关键字: \n" + ref.Keyword + "\n")
		buf.WriteString("标题: \n" + ref.Title + "\n")
		buf.WriteString("URL: \n" + ref.URL + "\n")
		buf.WriteString("正文内容: \n" + ref.Content + "\n")
		buf.WriteString(fmt.Sprintf("[参考资料 %d 结束]\n----------\n\n", index+1))
	}
	return buf.String()
}

type Agent struct {
	tavilyCli   *tools.TavilyClient
	askRunnable compose.Runnable[string, *askOutput]
}

func NewAgent() (*Agent, error) {
	a := &Agent{}
	ctx := context.Background()
	cfg := config.GetMainConfig()

	tavilyCli, err := tools.NewTavilyClient(ctx, &tools.TavilyConfig{
		APIKey: cfg.Tavily.APIKey,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create tavily client: %w", err)
	}
	a.tavilyCli = tavilyCli

	askGraph, err := a.createRunnableGraph(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create graph: %w", err)
	}
	// 编译 graph，将节点、边、分支转化为面向运行时的结构。由于 graph 中存在环，使用 AnyPredecessor 模式，同时设置运行时最大步数。
	askRunnable, err := askGraph.(*compose.Graph[string, *askOutput]).Compile(ctx,
		compose.WithNodeTriggerMode(compose.AnyPredecessor),
		compose.WithMaxRunSteps(100),
	)
	if err != nil {
		return nil, err
	}
	a.askRunnable = askRunnable
	return a, nil
}

func (a *Agent) Process(ctx context.Context, taskID string, initialMsg protocol.Message,
	handle taskmanager.TaskHandle) error {
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

	sr, err := a.askRunnable.Stream(ctx, part.Text,
		compose.WithCallbacks(callbackHandlers...))
	if err != nil {
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
	graph := compose.NewGraph[string, *askOutput](
		compose.WithGenLocalState(func(ctx context.Context) *askState {
			return &askState{}
		}),
	)

	startLambda := compose.InvokableLambda(
		func(ctx context.Context, input string) (output []*schema.Message, err error) {
			return []*schema.Message{}, nil
		},
	)
	_ = graph.AddLambdaNode("Lambda:start", startLambda,
		compose.WithNodeName("Lambda:start"),
		compose.WithStatePreHandler(
			func(ctx context.Context, in string, state *askState) (string, error) {
				state.UserInput = in
				return in, nil
			},
		),
	)

	thinkModel, err := deepseek.NewChatModel(ctx, &deepseek.ChatModelConfig{
		APIKey:  cfg.LLM.APIKey,
		BaseURL: cfg.LLM.URL,
		Model:   cfg.LLM.ReasoningModel,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create thinking model")
	}
	_ = graph.AddChatModelNode("ChatModel:think", thinkModel,
		compose.WithNodeName("ChatModel:think"),
		compose.WithStatePreHandler(
			func(ctx context.Context, in []*schema.Message, state *askState) ([]*schema.Message, error) {
				var out []*schema.Message
				out, err := AskThinkingModelSystemPrompt.Format(ctx, map[string]any{
					"meta_info": map[string]interface{}{
						"current_date": time.Now().Format("2006-01-02"),
					},
					"question":         state.UserInput,
					"max_search_words": defaultMaxSearchWords,
					"reference":        state.Reference.String(),
				})
				if err != nil {
					return nil, fmt.Errorf("failed to format system prompt, %w", err)
				}
				return out, nil
			},
		),
	)
	_ = graph.AddLambdaNode("ChatModel:think_to_list", compose.ToList[*schema.Message]())

	searchLambda := compose.InvokableLambda(
		func(ctx context.Context, input *schema.Message) ([]reference, error) {
			keywords := strings.Split(input.Content, ";")
			if err != nil {
				return nil, fmt.Errorf("failed to format prompt: %w", err)
			}
			var out []reference
			var mu sync.Mutex
			var handlers []func() error
			for index := range keywords {
				keyword := keywords[index]
				handlers = append(handlers, func() error {
					response, err := a.tavilyCli.Search(ctx, &tools.SearchRequest{
						Query:       keyword,
						SearchDepth: "basic",
						MaxResults:  defaultMaxResult,
					})
					if err != nil {
						log.ErrorContextf(ctx, "failed to search %s, err: %v", keyword, err)
						return nil
					}

					mu.Lock()
					defer mu.Unlock()
					for _, v := range response.Results {
						out = append(out, reference{
							Keyword: keyword,
							Title:   v.Title,
							URL:     v.URL,
							Content: v.Content,
						})
					}
					return nil
				})
			}
			if err = trpc.GoAndWait(handlers...); err != nil {
				return nil, fmt.Errorf("failed to search informations")
			}

			_ = compose.ProcessState(ctx, func(ctx context.Context, s *askState) error {
				s.Reference = append(s.Reference, out...)
				return nil
			})

			return out, nil
		},
	)
	_ = graph.AddLambdaNode("Lambda:search", searchLambda,
		compose.WithNodeName("Lambda:search"),
	)

	_ = graph.AddLambdaNode("Lambda:search_transform", compose.InvokableLambda(
		func(ctx context.Context, input []reference) (output []*schema.Message, err error) {
			return []*schema.Message{}, nil
		},
	))

	summaryModel, err := deepseek.NewChatModel(ctx, &deepseek.ChatModelConfig{
		APIKey:  cfg.LLM.APIKey,
		BaseURL: cfg.LLM.URL,
		Model:   cfg.LLM.ReasoningModel,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create thinking model")
	}
	_ = graph.AddChatModelNode("ChatModel:summary", summaryModel,
		compose.WithNodeName("ChatModel:summary"),
		compose.WithStatePreHandler(
			func(ctx context.Context, in []*schema.Message, state *askState) ([]*schema.Message, error) {
				var out []*schema.Message
				out, err := AskSummaryModelSystemPrompt.Format(ctx, map[string]any{
					"meta_info": map[string]interface{}{
						"current_date": time.Now().Format("2006-01-02"),
					},
					"reference": state.Reference.String(),
					"question":  state.UserInput,
				})
				if err != nil {
					return nil, fmt.Errorf("failed to format system prompt, %w", err)
				}
				return out, nil
			},
		),
	)

	_ = graph.AddLambdaNode("Lambda:output",
		compose.InvokableLambda(
			func(ctx context.Context, input *schema.Message) (output *askOutput, err error) {
				return &askOutput{Content: input.Content}, nil
			},
		),
		compose.WithStatePostHandler(
			func(ctx context.Context, out *askOutput, state *askState) (*askOutput, error) {
				out.Reference = state.Reference
				return out, nil
			},
		),
	)

	// 创建连线
	_ = graph.AddEdge(compose.START, "Lambda:start")
	_ = graph.AddEdge("Lambda:start", "ChatModel:think")

	_ = graph.AddBranch("ChatModel:think", compose.NewGraphBranch(
		func(ctx context.Context, in *schema.Message) (endNode string, err error) {
			var validateKeywords = func() bool {
				keywords := strings.Split(in.Content, ";")
				for _, keyword := range keywords {
					if len(keyword) > 100 {
						return false
					}
				}
				return true
			}
			if strings.Contains(in.Content, "无需检索") ||
				!validateKeywords() {
				return "ChatModel:think_to_list", nil
			}
			return "Lambda:search", nil
		}, map[string]bool{
			"ChatModel:think_to_list": true,
			"Lambda:search":           true,
		}))
	_ = graph.AddEdge("Lambda:search", "Lambda:search_transform")
	_ = graph.AddEdge("Lambda:search_transform", "ChatModel:think")
	_ = graph.AddEdge("ChatModel:think_to_list", "ChatModel:summary")
	_ = graph.AddEdge("ChatModel:summary", "Lambda:output")
	_ = graph.AddEdge("Lambda:output", compose.END)

	return graph, nil
}

type tag struct {
	id   string
	name string
}

func (t tag) String(entering bool) string {
	if entering {
		return fmt.Sprintf("<%s:%s>", t.name, t.id)
	}
	return fmt.Sprintf("</%s:%s>", t.name, t.id)
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
	case "ChatModel:think":
	case "Lambda:search":
		cb.updateWorkingTaskStatus(ctx, "\n>\n>\n> 🔍正在执行网络搜索…\n")
	case "ChatModel:summary":
		cb.updateWorkingTaskStatus(ctx, "\n>\n>\n")
	}
	return ctx
}

func (cb *callbackHandler) OnEnd(ctx context.Context, info *callbacks.RunInfo,
	output callbacks.CallbackOutput) context.Context {
	log.InfoContextf(ctx, "OnEnd: name=%s, type=%s, compoment=%s", info.Name, info.Type, info.Component)
	if cb.handle == nil {
		return ctx
	}

	switch info.Name {
	case "Lambda:search":
		var refNum int
		_ = compose.ProcessState(ctx, func(ctx context.Context, s *askState) error {
			refNum = len(s.Reference)
			return nil
		})
		cb.updateWorkingTaskStatus(ctx, fmt.Sprintf("> ✅网络搜索执行完毕，共搜到 %d 篇参考文献\n>\n>\n", refNum))
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
	case "ChatModel:think":
		return cb.processThinkNodeOutput(ctx, info, output)
	case "ChatModel:summary":
		return cb.processSummaryNodeOutput(ctx, info, output)
	}

	return ctx
}

func (cb *callbackHandler) processThinkNodeOutput(ctx context.Context, info *callbacks.RunInfo,
	output *schema.StreamReader[callbacks.CallbackOutput]) context.Context {

	cb.wg.Add(1)
	trpc.Go(ctx, time.Minute*10, func(ctx context.Context) {
		defer cb.wg.Done()
		defer output.Close() // remember to close the stream in defer

		var isReasoningContent bool

		for {
			frame, err := output.Recv()
			if errors.Is(err, io.EOF) {
				// finish
				break
			}
			if err != nil {
				log.ErrorContextf(ctx, "processThinkNodeOutput failed, err: %v", err)
				return
			}

			callbackOutput, ok := frame.(*model.CallbackOutput)
			if !ok {
				log.ErrorContextf(ctx, "invalid message content: %+v", frame)
				return
			}
			reasoningContent, ok := deepseek.GetReasoningContent(callbackOutput.Message)
			if ok {
				if !isReasoningContent {
					cb.updateWorkingTaskStatus(ctx, "> ")
				}
				isReasoningContent = true
				content := strings.ReplaceAll(reasoningContent, "\n", "\n> ")
				cb.updateWorkingTaskStatus(ctx, content)
			}
		}
	})

	return ctx
}

func (cb *callbackHandler) processSummaryNodeOutput(ctx context.Context, info *callbacks.RunInfo,
	output *schema.StreamReader[callbacks.CallbackOutput]) context.Context {
	cb.wg.Add(1)
	trpc.Go(ctx, time.Minute*10, func(ctx context.Context) {
		defer cb.wg.Done()
		defer output.Close() // remember to close the stream in defer

		var isReasoningContent bool

		for {
			frame, err := output.Recv()
			if errors.Is(err, io.EOF) {
				// finish
				break
			}
			if err != nil {
				log.ErrorContextf(ctx, "processThinkNodeOutput failed, err: %v", err)
				return
			}

			callbackOutput, ok := frame.(*model.CallbackOutput)
			if !ok {
				log.ErrorContextf(ctx, "invalid message content: %+v", frame)
				return
			}
			reasoningContent, ok := deepseek.GetReasoningContent(callbackOutput.Message)
			if ok {
				if !isReasoningContent {
					cb.updateWorkingTaskStatus(ctx, "> ")
				}
				isReasoningContent = true
				content := strings.ReplaceAll(reasoningContent, "\n", "\n> ")
				cb.updateWorkingTaskStatus(ctx, content)
				continue
			}

			if isReasoningContent {
				cb.updateWorkingTaskStatus(ctx, "\n")
				isReasoningContent = false
			}
			cb.updateWorkingTaskStatus(ctx, callbackOutput.Message.Content)
		}
	})

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
