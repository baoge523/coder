package deepseek

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino-ext/components/model/deepseek"
	"github.com/cloudwego/eino/components/model"
)

func NewChatModel() model.ToolCallingChatModel {
	ctx := context.Background()
	apiKey := "sk-77a82b7c3b7a4d0c97a4ca36aab90dce" // apikey

	// 创建 deepseek 模型
	cm, err := deepseek.NewChatModel(ctx, &deepseek.ChatModelConfig{
		APIKey: apiKey,
		Model:  "deepseek-reasoner", //   DeepSeek-V3-0324
		// BaseURL:     "https://api.deepseek.com/v1",
		MaxTokens:   5000,
		Temperature: 1.0,
	})

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return cm

}
