package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "projects/protobuf/usecase/generated_code/proto"
)

type server struct {
	pb.UnimplementedManagerServiceServer
}

func (s *server) CreateManager(ctx context.Context, req *pb.RequestManagerData) (*pb.Response, error) {
	reqStr, _ := json.Marshal(req)
	fmt.Printf("req body: %s \n", reqStr)
	return &pb.Response{Status: 200, Message: "ok"}, nil
}

func main() {
	// 通过net创建一个监听器
	listen, err := net.Listen("tcp", "127.0.0.1:2222")
	if err != nil {
		log.Fatal("net listen error")
	}
	// 获取grpc的server
	grpcServer := grpc.NewServer()
	// 将功能服务注册到 grpcServer中
	pb.RegisterManagerServiceServer(grpcServer, &server{})
	fmt.Printf("grpcServer.Serve launch \n")

	// 启动服务，监听端口
	if er := grpcServer.Serve(listen); er != nil {
		log.Fatal("grpcServer.Serve error")
	}
}
