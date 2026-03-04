# gRPC-Go Demo 项目

这是一个简单的 gRPC-Go 示例项目，包含客户端和服务端实现。

## 项目结构

```
grpc-demo-new/
├── proto/              # Protocol Buffers 定义
│   └── user.proto      # 用户服务 proto 文件
├── server/             # 服务端代码
│   └── server.go       # gRPC 服务端实现
├── client/             # 客户端代码
│   └── client.go       # gRPC 客户端实现
├── go.mod              # Go 模块文件
└── README.md           # 项目说明
```

## 功能说明

该项目实现了一个简单的用户服务，包含以下 RPC 方法：

- `GetUser`: 根据用户 ID 获取用户信息
- `CreateUser`: 创建新用户

## 使用步骤

### 1. 安装依赖

```bash
go mod tidy
```

### 2. 生成 gRPC 代码

首先需要安装 protoc 编译器和相关插件：

```bash
# 安装 protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# 安装 protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

然后生成代码：

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/user.proto
```

### 3. 启动服务端

```bash
go run server/server.go
```

服务端默认监听在 `localhost:50051`

### 4. 运行客户端

在另一个终端窗口运行：

```bash
go run client/client.go
```

## 特性

- 使用 Protocol Buffers 定义服务接口
- 实现了服务端和客户端拦截器（Interceptor）
- 支持 metadata 传递
- 支持超时控制
- 记录请求日志和耗时

## 拦截器功能

### 服务端拦截器
- 记录客户端 IP 地址
- 记录请求 metadata
- 记录方法调用耗时

### 客户端拦截器
- 自动添加客户端信息到 metadata
- 记录方法调用耗时

## 参考资料

- [gRPC-Go 官方文档](https://grpc.io/docs/languages/go/)
- [Protocol Buffers 文档](https://protobuf.dev/)
