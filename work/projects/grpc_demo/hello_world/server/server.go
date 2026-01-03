/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"

	"net"
	"projects/grpc_demo/hello_world"
	"projects/grpc_demo/hello_world/log"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	hello_world.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *hello_world.HelloRequest) (*hello_world.HelloReply, error) {
	select {
	case <-ctx.Done():
		log.InfoContextf(ctx, "ctx done")
		return nil, ctx.Err()
	case <-time.After(1 * time.Second):
		log.InfoContextf(ctx, "sleep 3 seconds")
	}
	log.InfoContextf(ctx, "Received: %s", in.GetName())
	log.Infof("Received: %s", in.GetName())
	return &hello_world.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer log.Sync()
	// add interceptor
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(metadataFilter))
	hello_world.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func metadataFilter(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

	var fields []log.Field
	if p, ok := peer.FromContext(ctx); ok {
		fields = append(fields, log.Field{
			Key:   "caller_ip",
			Value: p.Addr.String(),
		})
	}
	if incomingContext, ok := metadata.FromIncomingContext(ctx); ok {
		if incomingContext != nil && len(incomingContext.Get("caller_service")) > 0 {
			fields = append(fields, log.Field{"caller_service", incomingContext.Get("caller_service")[0]})
		}
		if incomingContext != nil && len(incomingContext.Get("caller_version")) > 0 {
			fields = append(fields, log.Field{"caller_version", incomingContext.Get("caller_version")[0]})
		}
	}
	service, method := parse(info.FullMethod)
	fields = append(fields, log.Field{"callee_service", service})
	fields = append(fields, log.Field{"callee_method", method})

	// log info
	ctx, _ = log.NewLogContext(ctx, fields...)
	return handler(ctx, req)
}

func parse(sm string) (string, string) {
	if sm != "" && sm[0] == '/' {
		sm = sm[1:]
	}
	pos := strings.LastIndex(sm, "/")
	if pos == -1 {
		return "", ""
	}
	service := sm[:pos]  // 服务名称
	method := sm[pos+1:] // 方法名称
	return service, method
}
