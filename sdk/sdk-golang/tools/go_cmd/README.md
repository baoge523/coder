
## go cmd

[go cmd](https://pkg.go.dev/cmd/go)
Go is a tool for managing Go source code.

Usage
```text
go <command> [arguments]
```
```text
bug         start a bug report
build       compile packages and dependencies
clean       remove object files and cached files
doc         show documentation for package or symbol
env         print Go environment information
fix         update packages to use new APIs
fmt         gofmt (reformat) package sources
generate    generate Go files by processing source
get         add dependencies to current module and install them
install     compile and install packages and dependencies
list        list packages or modules
mod         module maintenance
work        workspace maintenance
run         compile and run Go program
telemetry   manage telemetry data and settings
test        test packages
tool        run specified go tool
version     print Go version
vet         report likely mistakes in packages
```


### go clean
```text
这个清空模块信息的操作不用轻易执行，会情况所有的的依赖，如果你同时打开多个项目，那么这些项目都需要重新执行go mod tidy来下载对应的依赖信息
go clean -modcache 
```