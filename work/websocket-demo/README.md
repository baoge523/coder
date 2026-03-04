# WebSocket Header 验证项目

这个项目用于验证客户端在 WebSocket 升级握手时能否获取到服务端设置的自定义响应头。

## 项目结构

```
websocket-demo/
├── go.mod              # Go 模块文件
├── server/
│   └── main.go        # WebSocket 服务端
├── client/
│   └── main.go        # WebSocket 客户端
└── README.md          # 项目说明
```

## 核心功能

### 服务端
- 在 WebSocket 升级时设置自定义响应头：
  - `Authorization-Status: SUCCESS`
  - `Custom-Header: test-value`
  - `Server-Time: 1234567890`
- 接收并回显客户端消息

### 客户端
- 连接到 WebSocket 服务器
- 打印所有响应头信息
- 验证自定义响应头是否成功接收
- 支持交互式消息发送

## 使用方法

### 1. 安装依赖

```bash
cd work/websocket-demo
go mod tidy
```

### 2. 启动服务端

```bash
go run server/main.go
```

服务端将在 `localhost:8080` 启动，WebSocket 端点为 `/ws`

### 3. 启动客户端

在新的终端窗口中运行：

```bash
go run client/main.go
```

## 验证结果

客户端启动后会自动：
1. 连接到服务器
2. 打印所有响应头信息
3. 验证自定义响应头（Authorization-Status、Custom-Header、Server-Time）
4. 发送测试消息
5. 支持交互式输入（输入 'quit' 退出）

## 预期输出

客户端应该能看到类似以下输出：

```
✅ WebSocket 连接建立成功！

========== 服务器响应头信息 ==========
Upgrade: websocket
Connection: Upgrade
Sec-Websocket-Accept: ...
Authorization-Status: SUCCESS
Custom-Header: test-value
Server-Time: 1234567890

========== 验证自定义响应头 ==========
✅ Authorization-Status: SUCCESS
✅ Custom-Header: test-value
✅ Server-Time: 1234567890
```

## 技术要点

- 使用 `github.com/gorilla/websocket` 库
- 服务端通过 `upgrader.Upgrade(w, r, respHeader)` 的第三个参数传递响应头
- 客户端通过 `websocket.DefaultDialer.Dial()` 返回的 `*http.Response` 获取响应头
- 响应头格式为 `map[string][]string`
