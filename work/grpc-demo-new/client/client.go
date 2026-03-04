package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "grpc-demo-new/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	
	// 建立连接
	conn, err := grpc.NewClient(*addr, 
		grpc.WithUnaryInterceptor(clientInterceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	
	c := pb.NewUserServiceClient(conn)
	
	// 测试 CreateUser
	createUser(c)
	
	// 测试 GetUser
	getUser(c)
}

// createUser 创建用户
func createUser(c pb.UserServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	resp, err := c.CreateUser(ctx, &pb.CreateUserRequest{
		Username: "zhangsan",
		Email:    "zhangsan@example.com",
	})
	if err != nil {
		log.Fatalf("CreateUser failed: %v", err)
	}
	
	log.Printf("CreateUser response: user_id=%d, message=%s", resp.GetUserId(), resp.GetMessage())
}

// getUser 获取用户信息
func getUser(c pb.UserServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	resp, err := c.GetUser(ctx, &pb.GetUserRequest{
		UserId: 12345,
	})
	if err != nil {
		log.Fatalf("GetUser failed: %v", err)
	}
	
	log.Printf("GetUser response: user_id=%d, username=%s, email=%s, created_at=%d",
		resp.GetUserId(), resp.GetUsername(), resp.GetEmail(), resp.GetCreatedAt())
}

// clientInterceptor 客户端拦截器
func clientInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 添加 metadata
	md := metadata.Pairs(
		"client_name", "grpc-demo-client",
		"client_version", "1.0.0",
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	
	start := time.Now()
	
	// 调用远程方法
	err := invoker(ctx, method, req, reply, cc, opts...)
	
	log.Printf("Method: %s, Duration: %v, Error: %v", method, time.Since(start), err)
	
	return err
}
