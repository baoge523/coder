package memory

import (
	"context"

	"github.com/cloudwego/eino/schema"
)

const (
	StateKeyCurrentTaskID         = "sk:current_task_id"
	StateKeyCurrentAgentMessageID = "sk:current_agent_message_id"
	StateKeyCurrentUserEventID    = "sk:current_user_event_id"
)

type Factory interface {
	// Get 获取内存
	Get(ctx context.Context, userID string) (Memory, error)
}

type Memory interface {
	// GetUserID 获取用户ID
	GetUserID(ctx context.Context) string
	// GetState 获取状态
	GetState(ctx context.Context) (map[string]string, error)
	// SetState 设置状态
	SetState(ctx context.Context, fields ...string) error
	// GetConversation 获取会话
	GetConversation(ctx context.Context, id string) (Conversation, error)
	// GetCurrentConversation 获取当前会话
	GetCurrentConversation(ctx context.Context) (Conversation, error)
	// NewConversation 创建新会话
	NewConversation(ctx context.Context) (Conversation, error)
	// ListConversations 列出会话
	ListConversations(ctx context.Context) ([]string, error)
	// DeleteConversation 删除会话
	DeleteConversation(ctx context.Context, id string) error
}

type Conversation interface {
	// GetMemory 获取用户Memory
	GetMemory(ctx context.Context) Memory
	// GetID 获取会话ID
	GetID(ctx context.Context) string
	// Append 追加消息，如果msg的ID一致，那么更新消息
	Append(ctx context.Context, msgID string, msg *schema.Message) error
	// GetMessages 获取消息
	GetMessages(ctx context.Context) ([]*schema.Message, error)
	// GetMessage 获取单条消息
	GetMessage(ctx context.Context, msgID string) (*schema.Message, error)
}
