package deepseek

import (
	"context"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"time"
)

/***

这里的消息模版信息很重要，就是告诉ai以什么样的方式去分析对应的效果

*/

// 创建模板，使用 FString 格式
var template = prompt.FromMessages(schema.GoTemplate,
	// 系统消息模板
	schema.SystemMessage("你是一个{{.role}}。你需要用{{.style}}的语气回答问题。你的目标是帮助程序员保持积极乐观的心态，提供技术建议的同时也要关注他们的心理健康。"),

	// 插入需要的对话历史（新对话的话这里不填）
	schema.MessagesPlaceholder("chat_history", true),

	// 用户消息模板
	schema.UserMessage("问题: {{.question}}"),
)

// 使用模板生成消息
var messages, err = template.Format(context.Background(), map[string]any{
	"role":     "程序员鼓励师",
	"style":    "积极、温暖且专业",
	"question": "我的代码一直报错，感觉好沮丧，该怎么办？",
	// 对话历史（这个例子里模拟两轮对话历史）
	"chat_history": []*schema.Message{
		schema.UserMessage("你好"),
		schema.AssistantMessage("嘿！我是你的程序员鼓励师！记住，每个优秀的程序员都是从 Debug 中成长起来的。有什么我可以帮你的吗？", nil),
		schema.UserMessage("我觉得自己写的代码太烂了"),
		schema.AssistantMessage("每个程序员都经历过这个阶段！重要的是你在不断学习和进步。让我们一起看看代码，我相信通过重构和优化，它会变得更好。记住，Rome wasn't built in a day，代码质量是通过持续改进来提升的。", nil),
	},
})

// ---------

// 创建模板，使用 FString 格式
var ampTemp = prompt.FromMessages(schema.FString,
	// 系统消息模板
	schema.SystemMessage(""),

	// 用户消息模板
	schema.UserMessage("当前时间: {current_time}\n    要分析的告警历史的明细信息: {current_alarm_history}\n    要分析的告警历史所关联的告警策略: {policy_info}\n\n    与当前要分析的告警历史是同一个告警策略的其他近期的告警历史(注意, 总数是{policy_related_total}, 这里只返回近期的 {policy_related_alarm_history_size} 条告警历史): \n    <policy_related_alarm_history>\n    {policy_related_alarm_history}\n    </policy_related_alarm_history>\n\n    与当前要分析的告警历史属于同一个告警对象的其他近期的告警历史:\n    <alarm_object_related_alarm_history>\n    {alarm_object_related_alarm_history}\n    </alarm_object_related_alarm_history>"),
)

// 使用模板生成消息
var ampMessage, ampErr = ampTemp.Format(context.Background(), map[string]any{
	"current_time":                       time.Now().Format(time.RFC3339),
	"policy_info":                        "",
	"current_alarm_history":              "",
	"policy_related_alarm_history":       "",
	"policy_related_alarm_history_size":  14,
	"policy_related_total":               173,
	"alarm_object_related_alarm_history": "[]",
})
