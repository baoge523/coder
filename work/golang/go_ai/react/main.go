package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/eino-ext/components/model/openai"
	einomcp "github.com/cloudwego/eino-ext/components/tool/mcp"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
)

var (
	amapURL = "https://mcp.amap.com/sse?key=xxxx" // 地图提供的mcp server访问地址
	apiKey  = "xxxx"                              // 模型apiKey
	baseURL = "xxxx"                              //
	model   = "xxxx"
)

func connectMCP(ctx context.Context, serverURL string) ([]tool.BaseTool, []*schema.ToolInfo, error) {
	// MCP 客户端连接 MCP server端
	cli, err := client.NewSSEMCPClient(serverURL)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect: %w", err)
	}
	err = cli.Start(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to restart: %w", err)
	}
	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "go-client",
		Version: "1.0.0",
	}
	_, err = cli.Initialize(ctx, initRequest)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize: %w", err)
	}
	// 通过einomcp获取mcp的tools
	tools, err := einomcp.GetTools(ctx, &einomcp.Config{Cli: cli})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get tools: %w", err)
	}

	// 从工具中获取工具详情信息
	var toolsInfo []*schema.ToolInfo
	for _, v := range tools {
		toolInfo, err := v.Info(ctx)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get tool info: %w", err)
		}
		toolsInfo = append(toolsInfo, toolInfo)
	}
	// 返回工具信息和工具详情
	return tools, toolsInfo, nil
}

// state 本地状态
// 将模型输入、输出的上下文记录到全局状态中完成ReAct的迭代
type state struct {
	Messages []*schema.Message
}

/**
  这里的用例主要体现了单agent的场景

   agent通过mcp client 访问 mcp server获取tool，在意向分析中，判断是否需要调用工具从而判断是否需要结束

*/

func main() {
	ctx := context.Background()

	// 创建一个图
	g := compose.NewGraph[map[string]any, *schema.Message](
		// 生成本地状态
		compose.WithGenLocalState(func(ctx context.Context) *state {
			return &state{}
		}))

	// 组装提示词
	promptTemplate := prompt.FromMessages(
		schema.FString,                                              // 占位符格式
		schema.SystemMessage("你是一个智能助手，请帮我解决以下问题"), // 系统的prompt
		schema.UserMessage("{location}今天天气怎么样？"),             // 用户的prompt
	)
	// 连接指定的mcp server获取工具信息
	tools, toolsInfo, err := connectMCP(ctx, amapURL)
	if err != nil {
		panic(err)
	}

	// 创建LLM
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  apiKey,
		BaseURL: baseURL,
		Model:   model,
	})
	if err != nil {
		panic(err)
	}
	// LLM 绑定工具；大模型绑定工具获取Function Call能力
	if err = chatModel.BindTools(toolsInfo); err != nil {
		panic(err)
	}
	toolsNode, err := compose.NewToolNode(ctx, &compose.ToolsNodeConfig{
		Tools: tools,
	})

	// 给图创建节点信息
	_ = g.AddChatTemplateNode("ChatTemplate", promptTemplate)
	_ = g.AddChatModelNode("ChatModel", chatModel,


		// 记录用户历史提问
		compose.WithStatePreHandler(
			func(ctx context.Context, in []*schema.Message, state *state) ([]*schema.Message, error) {
				state.Messages = append(state.Messages, in...)
				return state.Messages, nil
			},
		),
		// 记录用户历史提供的AI回答
		compose.WithStatePostHandler(
			func(ctx context.Context, out *schema.Message, state *state) (*schema.Message, error) {
				state.Messages = append(state.Messages, out)
				return out, nil
			},
		),
	)
	// 工具Node
	_ = g.AddToolsNode("ToolNode", toolsNode)

	// 图的边 compose.START表示起点，compose.END表示终点
	_ = g.AddEdge(compose.START, "ChatTemplate")
	_ = g.AddEdge("ChatTemplate", "ChatModel")

	// AddBranch 表示 ChatModel节点后，有分支判断，
	_ = g.AddBranch("ChatModel", compose.NewGraphBranch(

		// 具体的分支判断,返回判断流程走到哪个节点，如果有错误，直接终止流程
		func(ctx context.Context, in *schema.Message) (endNode string, err error) {
			if len(in.ToolCalls) == 0 {
				return compose.END, nil
			}
			return "ToolNode", nil
		}, map[string]bool{
			"ToolNode":  true,
			compose.END: true,
		}))
	_ = g.AddEdge("ToolNode", "ChatModel")

	// 编译图，并设置最大的运行步骤为10步，避免一直循环判断
	r, err := g.Compile(ctx, compose.WithMaxRunSteps(10))
	if err != nil {
		panic(err)
	}

	in := map[string]any{"location": "广州"}

	// 执行，传入 入参
	ret, err := r.Invoke(ctx, in)
	if err != nil {
		panic(err)
	}
	log.Println("invoke result: ", ret)
}
