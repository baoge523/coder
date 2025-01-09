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