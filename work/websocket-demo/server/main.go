package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("收到 WebSocket 升级请求，来自: %s", r.RemoteAddr)

	// 设置响应头，这些头会在升级握手时返回给客户端
	respHeader := make(map[string][]string)
	respHeader["Authorization-Status"] = []string{"SUCCESS"}
	respHeader["Custom-Header"] = []string{"test-value"}
	respHeader["Server-Time"] = []string{fmt.Sprintf("%d", 1234567890)}

	// 升级 HTTP 连接为 WebSocket
	conn, err := upgrader.Upgrade(w, r, respHeader)
	if err != nil {
		log.Printf("WebSocket 升级失败: %v", err)
		return
	}
	defer conn.Close()

	log.Printf("WebSocket 连接建立成功，客户端: %s", conn.RemoteAddr())

	// 发送欢迎消息
	err = conn.WriteMessage(websocket.TextMessage, []byte("欢迎连接到 WebSocket 服务器！"))
	if err != nil {
		log.Printf("发送消息失败: %v", err)
		return
	}

	// 循环读取客户端消息
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("读取消息失败: %v", err)
			break
		}

		log.Printf("收到消息 [类型: %d]: %s", messageType, string(message))

		// 回显消息
		response := fmt.Sprintf("服务器收到: %s", string(message))
		err = conn.WriteMessage(messageType, []byte(response))
		if err != nil {
			log.Printf("发送响应失败: %v", err)
			break
		}
	}

	log.Printf("WebSocket 连接关闭，客户端: %s", conn.RemoteAddr())
}

func main() {
	http.HandleFunc("/ws", wsHandler)

	addr := ":8080"
	log.Printf("WebSocket 服务器启动，监听地址: %s", addr)
	log.Printf("WebSocket 端点: ws://localhost%s/ws", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
