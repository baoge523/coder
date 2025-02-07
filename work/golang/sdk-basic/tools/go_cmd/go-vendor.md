## vendor

vendor主要是用于在没有网络时编译出go可执行文件

### build时，使用vendor
```linux
go build -mod=vendor
```

### 通过modules 生成 vendor
```text
go mod vendor
```
但是 go mod vendor 不会将modules中的cgo依赖拷贝到vendor中，所以需要使用工具将这些文件拷贝到vendor中

https://github.com/goware/modvendor

```text
go get -u github.com/goware/modvendor

GO111MODULE=on go mod vendor

modvendor -copy="**/*.c **/*.h **/*.proto" -v


tce的vender

modvendor -copy="**/*.c **/*.h **/*.proto **/*.h **/*.a **/*.dll" -v
```