package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/google/uuid"
)

var (
	apiKey    = "aaaaaa"
	baseURL   = "http://xxxx"
	modelName = "aaaaaa"
)

func main() {

	// 注册一个序列化类型，主要用于存放CheckPoint信息
	_ = compose.RegisterSerializableType[StateInfo]("state")

	ctx := context.Background()

	g := compose.NewGraph[map[string]any, *schema.Message](
		compose.WithGenLocalState(
			func(ctx context.Context) *StateInfo {
				return &StateInfo{}
			},
		),
	)
	_ = g.AddChatTemplateNode("ChatTemplate", newChatTemplate(ctx))
	_ = g.AddChatModelNode("ChatModel", newChatModel(ctx),
		compose.WithStatePreHandler(
			func(ctx context.Context, in []*schema.Message, state *StateInfo) ([]*schema.Message, error) {
				state.History = append(state.History, in...)
				return state.History, nil
			}),
		compose.WithStatePostHandler(
			func(ctx context.Context, out *schema.Message, state *StateInfo) (*schema.Message, error) {
				state.History = append(state.History, out)
				return out, nil
			}),
	)
	// 添加一个自定义的 LambdaNode； 这里使用的是 invoke类型的lambda
	_ = g.AddLambdaNode("HumanInTheLoop", compose.InvokableLambda(
		func(ctx context.Context, input *schema.Message) (output *schema.Message, err error) {
			var userInput string
			// 执行状态的回调函数，获取状态中的信息
			_ = compose.ProcessState(ctx, func(ctx context.Context, s *StateInfo) error {
				userInput = s.Input
				return nil
			})
			// 判断是否有用户输入数据，如果没有，直接中断
			if userInput == "" {
				return nil, compose.InterruptAndRerun
			}
			// 如果是n表示取消
			if strings.ToLower(userInput) == "n" {
				return nil, fmt.Errorf("user cancel")
			}
			return input, nil
		}))
	_ = g.AddToolsNode("ToolsNode", newToolsNode(ctx),
		compose.WithStatePreHandler(
			func(ctx context.Context, in *schema.Message, state *StateInfo) (*schema.Message, error) {
				return state.History[len(state.History)-1], nil
			}))
	_ = g.AddEdge(compose.START, "ChatTemplate")
	_ = g.AddEdge("ChatTemplate", "ChatModel")
	_ = g.AddEdge("ToolsNode", "ChatModel")
	_ = g.AddBranch("ChatModel",
		compose.NewGraphBranch(func(ctx context.Context, in *schema.Message) (endNode string, err error) {
			if len(in.ToolCalls) > 0 {
				return "HumanInTheLoop", nil
			}
			return compose.END, nil
		}, map[string]bool{"HumanInTheLoop": true, compose.END: true}))
	_ = g.AddEdge("HumanInTheLoop", "ToolsNode")

	// 会将checkpoint信息保存起来，供下次恢复使用
	runner, err := g.Compile(ctx, compose.WithCheckPointStore(newCheckPointStore(ctx)))
	if err != nil {
		panic(err)
	}

	taskID := uuid.New().String()
	var history []*schema.Message
	var userInput string
	for {
		result, err := runner.Invoke(ctx, map[string]any{"name": "Alice", "location": "广州"},
			// 检查点，用于暂存信息
			compose.WithCheckPointID(taskID),
			// 修改状态信息
			compose.WithStateModifier(func(ctx context.Context, path compose.NodePath, state any) error {
				state.(*StateInfo).Input = userInput
				state.(*StateInfo).History = history
				return nil
			}))
		if err == nil {
			fmt.Printf("执行成功: %s", result.Content)
			break
		}
		// 如果是中断错误，表示都没有输入，等待用户输入
		info, ok := compose.ExtractInterruptInfo(err)
		if !ok {
			log.Fatal(err)
		}
		history = info.State.(*StateInfo).History
		for _, tc := range history[len(history)-1].ToolCalls {
			fmt.Printf("使用工具: %s, 参数: %s\n请确认参数是否正确? (y/n): ", tc.Function.Name, tc.Function.Arguments)
			fmt.Scanln(&userInput)
		}
	}
}

// newChatTemplate 创建prompt
func newChatTemplate(_ context.Context) prompt.ChatTemplate {
	return prompt.FromMessages(schema.FString,
		schema.SystemMessage("You are a helpful assistant. If the user asks about the booking, call the \"BookTicket\" tool to book ticket. 请使用中文输出。"),
		schema.UserMessage("I'm {name}. Help me book a ticket to {location}"),
	)
}

// newChatModel 创建LLM
func newChatModel(ctx context.Context) *openai.ChatModel {
	cm, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  apiKey,
		BaseURL: baseURL,
		Model:   modelName,
	})
	if err != nil {
		log.Fatal(err)
	}

	tools := getTools()
	var toolsInfo []*schema.ToolInfo
	for _, t := range tools {
		info, err := t.Info(ctx)
		if err != nil {
			log.Fatal(err)
		}
		toolsInfo = append(toolsInfo, info)
	}

	err = cm.BindTools(toolsInfo)
	if err != nil {
		log.Fatal(err)
	}
	return cm
}

type bookInput struct {
	Location             string `json:"location"`
	PassengerName        string `json:"passenger_name"`
	PassengerPhoneNumber string `json:"passenger_phone_number"`
}

// newToolsNode 获取工具
func newToolsNode(ctx context.Context) *compose.ToolsNode {
	tools := getTools()

	tn, err := compose.NewToolNode(ctx, &compose.ToolsNodeConfig{Tools: tools})
	if err != nil {
		log.Fatal(err)
	}
	return tn
}

func newCheckPointStore(ctx context.Context) compose.CheckPointStore {
	return &StoreInfo{buf: make(map[string][]byte)}
}

type StateInfo struct {
	Input   string
	History []*schema.Message
}

func getTools() []tool.BaseTool {
	getWeather, err := utils.InferTool("BookTicket", "this tool can book ticket of the specific location",
		func(ctx context.Context, input bookInput) (output string, err error) {
			return "success", nil
		})
	if err != nil {
		fmt.Printf("get BookTicket error %v\n", err)
		return nil
	}

	return []tool.BaseTool{
		getWeather,
	}
}

type StoreInfo struct {
	buf map[string][]byte
}

func (m *StoreInfo) Get(_ context.Context, checkPointID string) ([]byte, bool, error) {
	data, ok := m.buf[checkPointID]
	return data, ok, nil
}

func (m *StoreInfo) Set(_ context.Context, checkPointID string, checkPoint []byte) error {
	m.buf[checkPointID] = checkPoint
	return nil
}
