.PHONY: proto server client clean

# 生成 protobuf 代码
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/user.proto

# 运行服务端
server:
	go run server/server.go

# 运行客户端
client:
	go run client/client.go

# 清理生成的文件
clean:
	rm -f proto/*.pb.go

# 安装依赖
deps:
	go mod tidy
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
