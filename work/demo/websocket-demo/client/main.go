package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	url := "ws://localhost:8080/ws"
	log.Printf("正在连接到 WebSocket 服务器: %s", url)

	// 创建 HTTP 请求头
	requestHeader := http.Header{}
	requestHeader.Add("User-Agent", "WebSocket-Client/1.0")

	// 连接到 WebSocket 服务器
	conn, resp, err := websocket.DefaultDialer.Dial(url, requestHeader)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	log.Println("✅ WebSocket 连接建立成功！")
	log.Println("\n========== 服务器响应头信息 ==========")

	// 打印所有响应头
	for key, values := range resp.Header {
		for _, value := range values {
			log.Printf("%s: %s", key, value)
		}
	}

	// 重点验证自定义响应头
	log.Println("\n========== 验证自定义响应头 ==========")
	authStatus := resp.Header.Get("Authorization-Status")
	if authStatus != "" {
		log.Printf("✅ Authorization-Status: %s", authStatus)
	} else {
		log.Println("❌ 未获取到 Authorization-Status 头")
	}

	customHeader := resp.Header.Get("Custom-Header")
	if customHeader != "" {
		log.Printf("✅ Custom-Header: %s", customHeader)
	} else {
		log.Println("❌ 未获取到 Custom-Header 头")
	}

	serverTime := resp.Header.Get("Server-Time")
	if serverTime != "" {
		log.Printf("✅ Server-Time: %s", serverTime)
	} else {
		log.Println("❌ 未获取到 Server-Time 头")
	}

	log.Println("\n========== 开始消息交互 ==========")

	// 启动消息接收协程
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("读取消息失败: %v", err)
				return
			}
			log.Printf("📩 收到服务器消息: %s", string(message))
		}
	}()

	// 启动用户输入协程
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		log.Println("请输入消息（输入 'quit' 退出）:")
		for scanner.Scan() {
			text := scanner.Text()
			if text == "quit" {
				log.Println("准备退出...")
				interrupt <- os.Interrupt
				return
			}

			err := conn.WriteMessage(websocket.TextMessage, []byte(text))
			if err != nil {
				log.Printf("发送消息失败: %v", err)
				return
			}
		}
	}()

	// 发送测试消息
	time.Sleep(100 * time.Millisecond)
	testMsg := "Hello, WebSocket Server!"
	log.Printf("📤 发送测试消息: %s", testMsg)
	err = conn.WriteMessage(websocket.TextMessage, []byte(testMsg))
	if err != nil {
		log.Printf("发送测试消息失败: %v", err)
	}

	// 等待中断信号
	select {
	case <-done:
		log.Println("连接已关闭")
	case <-interrupt:
		log.Println("收到中断信号，正在关闭连接...")

		// 发送关闭消息
		err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Printf("发送关闭消息失败: %v", err)
		}

		select {
		case <-done:
		case <-time.After(time.Second):
		}
	}

	fmt.Println("客户端已退出")
}
