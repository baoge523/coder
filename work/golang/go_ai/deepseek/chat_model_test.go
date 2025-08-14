package deepseek

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino/schema"
	"io"
	"log"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	ctx := context.Background()
	chatModel := NewChatModel()
	// 非流式响应
	//  	_, err := chatModel.Generate(ctx, ampMessage)
	// 流式响应
	fmt.Println(chatModel)
	streamResult, err := chatModel.Stream(ctx, ampMessage)
	if err != nil {
		fmt.Println(err)
		return
	}
	msgChan := reportStream2(streamResult)

	fullMsg := strings.Builder{}
	for msg := range msgChan {
		fullMsg.WriteString(msg)
	}
	fmt.Println(fullMsg.String())

}

// reportStream 可以将从sr中读取到的数据写入chan中，然后在chan中消费，但是读取完毕后，需要关闭chan
func reportStream(sr *schema.StreamReader[*schema.Message]) {
	defer sr.Close()

	fmt.Println(sr)
	i := 0
	for {
		message, err := sr.Recv()
		if err == io.EOF { // 流式输出结束
			return
		}
		if err != nil {
			log.Fatalf("recv failed: %v", err)
		}
		log.Printf("message[%d]: %+v\n", i, message)
		i++
	}
}
func reportStream2(sr *schema.StreamReader[*schema.Message]) chan string {
	msg := make(chan string, 1)
	go func() {
		i := 0
		defer sr.Close()
		defer close(msg)
		for {
			message, err := sr.Recv()
			if err == io.EOF { // 流式输出结束
				return
			}
			if err != nil {
				log.Fatalf("recv failed: %v", err)
			}
			msg <- message.Content
			i++
		}
	}()

	return msg
}
