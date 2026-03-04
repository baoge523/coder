package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "grpc-demo-new/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server 实现 UserService 服务
type server struct {
	pb.UnimplementedUserServiceServer
}

// GetUser 获取用户信息
func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Printf("Received GetUser request for user_id: %d", req.GetUserId())
	
	// 模拟数据库查询
	return &pb.GetUserResponse{
		UserId:    req.GetUserId(),
		Username:  fmt.Sprintf("user_%d", req.GetUserId()),
		Email:     fmt.Sprintf("user_%d@example.com", req.GetUserId()),
		CreatedAt: time.Now().Unix(),
	}, nil
}

// CreateUser 创建用户
func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Printf("Received CreateUser request: username=%s, email=%s", req.GetUsername(), req.GetEmail())
	
	// 模拟生成用户ID
	userId := time.Now().Unix()
	
	return &pb.CreateUserResponse{
		UserId:  userId,
		Message: fmt.Sprintf("User %s created successfully", req.GetUsername()),
	}, nil
}

func main() {
	flag.Parse()
	
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	// 创建 gRPC 服务器，添加拦截器
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(loggingInterceptor))
	pb.RegisterUserServiceServer(s, &server{})
	
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// loggingInterceptor 日志拦截器
func loggingInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()
	
	// 获取客户端信息
	if p, ok := peer.FromContext(ctx); ok {
		log.Printf("Client IP: %s", p.Addr.String())
	}
	
	// 获取 metadata
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("Metadata: %v", md)
	}
	
	// 调用实际的处理函数
	resp, err := handler(ctx, req)
	
	// 记录耗时
	log.Printf("Method: %s, Duration: %v, Error: %v", info.FullMethod, time.Since(start), err)
	
	return resp, err
}
