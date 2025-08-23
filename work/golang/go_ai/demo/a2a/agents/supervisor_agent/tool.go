package supervisor_agent

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/google/uuid"
	"trpc.group/trpc-go/trpc-a2a-go/protocol"
	"trpc.group/trpc-go/trpc-go/log"
)

type SendTask struct {
	agent *Agent
}

func (t *SendTask) Info(ctx context.Context) (*schema.ToolInfo, error) {
	var agentNameEnum []string
	for agentName := range t.agent.agentClientMap {
		agentNameEnum = append(agentNameEnum, agentName)
	}
	// 定义工具的信息
	return &schema.ToolInfo{
		Name: "send_task",
		Desc: "给Agent发送任务",
		// 定义工具的参数信息
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			// 指定调用哪个agent
			"agent_name": {
				Type:     "string",
				Desc:     "发送任务的Agent名字",
				Enum:     agentNameEnum,
				Required: true,
			},
			// 消息信息
			"message": {
				Type:     "string",
				Desc:     "发送的任务参数",
				Required: true,
			},
		}),
	}, nil
}

// SendTaskParams 参数定义
type SendTaskParams struct {
	AgentName string `json:"agent_name"`
	Message   string `json:"message"`
}

func (t *SendTask) StreamableRun(ctx context.Context, argumentsInJSON string,
	opts ...tool.Option) (*schema.StreamReader[string], error) {
	params := &SendTaskParams{}
	err := json.Unmarshal([]byte(argumentsInJSON), params)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal arguments: %w", err)
	}

	// 拿到目标agent
	clientAgent, ok := t.agent.agentClientMap[params.AgentName]
	if !ok {
		return nil, fmt.Errorf("agent %s not found", params.AgentName)
	}

	// 构造任务id，包含唯一id和agent名称，可能还需要包含其他的东西，方便最终问题
	taskID := model.TaskID{
		AgentName: params.AgentName,
		ID:        uuid.New().String(),
	}
	// 处理状态信息
	err = compose.ProcessState(ctx, func(ctx context.Context, s *state) error {
		// 设置任务id到内存状态中
		return s.mem.SetState(ctx, memory.StateKeyCurrentTaskID, taskID.Encode())
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update user state: %w", err)
	}
	// 通过A2A协议调用其他的agent
	taskChan, err := clientAgent.a2aClient.StreamTask(ctx, protocol.SendTaskParams{
		ID: taskID.Encode(),
		Message: protocol.Message{
			Role:  protocol.MessageRoleUser,
			Parts: []protocol.Part{protocol.NewTextPart(params.Message)},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to send task: %w", err)
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
}

type ClearMemory struct {
	agent *Agent
}

func (t *ClearMemory) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "clear_memory",
		Desc: "清除上下文，清除用户记忆",
		ParamsOneOf: schema.NewParamsOneOfByOpenAPIV3(&openapi3.Schema{
			Type: openapi3.TypeObject,
			Properties: map[string]*openapi3.SchemaRef{
				"all": {
					Value: &openapi3.Schema{
						Description: "always is True",
						Type:        openapi3.TypeBoolean,
						Enum: []interface{}{
							true,
						},
					},
				},
			},
			Required: []string{"all"},
		}),
	}, nil
}

func (t *ClearMemory) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	err := compose.ProcessState(ctx, func(ctx context.Context, s *state) error {
		if _, err := s.mem.NewConversation(ctx); err != nil {
			return fmt.Errorf("failed to new conversation: %w", err)
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to clear memory: %w", err)
	}
	log.InfoContextf(ctx, "clear memofy success")
	return "Success", nil
}
