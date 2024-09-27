
## go cmd

[do cmd](https://pkg.go.dev/cmd/go)


### go clean
```text
这个清空模块信息的操作不用轻易执行，会情况所有的的依赖，如果你同时打开多个项目，那么这些项目都需要重新执行go mod tidy来下载对应的依赖信息
go clean -modcache 
```