package lbshelper

import (
	"context"
	"fmt"

	"trpc.group/trpc-go/trpc-a2a-go/protocol"
	"trpc.group/trpc-go/trpc-a2a-go/server"
	"trpc.group/trpc-go/trpc-a2a-go/taskmanager"
	redistaskmanager "trpc.group/trpc-go/trpc-a2a-go/taskmanager/redis"
	tgoredis "trpc.group/trpc-go/trpc-database/goredis"
	"trpc.group/trpc-go/trpc-go/log"
)

func NewA2AServer(agent *Agent) (*server.A2AServer, error) {
	var err error
	agentCard := getAgentCard()
	processor := &TaskProcessor{}
	processor.agent = agent
	redisCli, err := tgoredis.New("trpc.redis.lbshelper")
	if err != nil {
		return nil, fmt.Errorf("failed to create redis client: %w", err)
	}
	taskManager, err := redistaskmanager.NewRedisTaskManager(redisCli, processor)
	if err != nil {
		return nil, fmt.Errorf("failed to create task manager: %w", err)
	}
	srv, err := server.NewA2AServer(agentCard, taskManager)
	if err != nil {
		return nil, fmt.Errorf("failed to create A2A server: %w", err)
	}
	return srv, nil
}

// Helper function to create a string pointer
func stringPtr(s string) *string {
	return &s
}

func getAgentCard() server.AgentCard {
	agentCard := server.AgentCard{
		Name:        "lbshelper",
		Description: stringPtr("一个行程智能助手，可以利用工具进行网络搜索，帮助用户解决行程规划相关的问题。"),
		Version:     "0.0.1",
		Provider: &server.AgentProvider{
			Organization: "a2a_samples",
		},
		Capabilities: server.AgentCapabilities{
			PushNotifications:      true, // Enable push notifications
			StateTransitionHistory: true, // MemoryTaskManager stores history
		},
		// Support text input/output
		DefaultInputModes:  []string{string(protocol.PartTypeText)},
		DefaultOutputModes: []string{string(protocol.PartTypeText)},
		Skills: []server.AgentSkill{
			{
				ID:          "query_weather",
				Name:        "根据城市名称或者标准adcode查询指定城市的天气",
				Description: stringPtr(""),
				Tags:        []string{"weather"},
				Examples:    []string{"深圳今天的天气怎么样"},
				InputModes:  []string{"text"},
				OutputModes: []string{"text"},
			},
			{
				ID:          "maps_direction",
				Name:        "查询城市之间的路线",
				Description: stringPtr(""),
				Tags:        []string{"weather"},
				Examples:    []string{"深圳到广州怎么走"},
				InputModes:  []string{"text"},
				OutputModes: []string{"text"},
			},
			{
				ID:          "travel_planning",
				Name:        "查询旅行攻略",
				Description: stringPtr("搜索目的地，查询旅行攻略"),
				Tags:        []string{"travel"},
				Examples:    []string{"广州长隆有哪些好玩的项目"},
				InputModes:  []string{"text"},
				OutputModes: []string{"text"},
			},
		},
	}
	return agentCard
}

type TaskProcessor struct {
	agent *Agent
}

func (t *TaskProcessor) Process(ctx context.Context, taskID string, initialMsg protocol.Message,
	handle taskmanager.TaskHandle) error {
	if err := t.agent.Process(ctx, taskID, initialMsg, handle); err != nil {
		log.ErrorContextf(ctx, "process %s fail, err: %v", initialMsg, err)
	}
	return nil
}
