## mac 相关的命令 以及环境变量信息

look for someone command bin path
```bash
which protoc
```

look for all bin path in mac
```bash
echo $path
```


### how to install protoc in my mac
```text
protoc 的版本是：libprotoc 3.21.12
command: brew install protobuf@21

protoc-gen-go 的版本是:  protoc-gen-go v1.28.0    
go evn command: go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0

protoc-gen-go-grpc 的版本是:  protoc-gen-go-grpc v1.2.0  
go evn command: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

-- those command solve why generate code whit `Any` types --

```


mac上面查看哪些服务会被自动拉起
```text
在 macOS 上，自动启动服务的配置通常位于以下位置：
1. **用户级服务**：
   - `~/Library/LaunchAgents/`：用户登录时启动的服务。
2. **系统级服务**：
   - `/Library/LaunchAgents/`：所有用户登录时启动的服务。
   - `/Library/LaunchDaemons/`：系统启动时启动的服务（无需用户登录）。
这些目录中的 `.plist` 文件定义了服务的启动行为和配置。你可以编辑或删除相应的 `.plist` 文件来修改自动启动的设置。
```