package supervisor_agent

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
	redisCli, err := tgoredis.New("trpc.redis.supervisor_agent")
	if err != nil {
		return nil, fmt.Errorf("failed to create redis client: %w", err)
	}
	// 生成任务处理器对象，其中生成的子任务会通过processor执行
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
		Name:        "supervisor_agent",
		Description: stringPtr(""),
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
		Skills:             []server.AgentSkill{},
	}
	return agentCard
}

type TaskProcessor struct {
	agent *Agent
}

// Process ai拆解后的子任务
func (t *TaskProcessor) Process(ctx context.Context, taskID string, initialMsg protocol.Message,
	handle taskmanager.TaskHandle) error {
	// 启动Agent任务
	if err := t.agent.Process(ctx, taskID, initialMsg, handle); err != nil {
		log.ErrorContextf(ctx, "process %s fail, err: %v", initialMsg, err)
	}
	return nil
}
