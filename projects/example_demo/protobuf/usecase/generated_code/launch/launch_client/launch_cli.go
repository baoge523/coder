package main

import (
	"context"
	"log"
	"projects/example_demo/protobuf/usecase/generated_code/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 通过grpc 获取客户端连接 -- 这个连接应该被复用 （连接池）
	conn, err := grpc.NewClient("127.0.0.1:2222", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("client error")
	}
	// 连接是一种资源，不用是需要关闭
	defer conn.Close()
	// pb通过连接创建客户端
	client := manager.NewManagerServiceClient(conn)

	//设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := client.CreateManager(ctx, &manager.RequestManagerData{
		Manager: &manager.Manager{
			Name: "zhangsan",
			Age:  18,
		},
	})
	if err != nil {
		log.Fatal("client CreateManager error")
	}

	log.Printf("message: %s", r.GetMessage())
}
