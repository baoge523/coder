## go mod 小工具
go mod 用于管理golang中的模块信息，我们可以通过go help mod 查看帮助文档
```linux
go help mod
```
```text
Usage:

        go mod <command> [arguments]

The commands are:

        download    download modules to local cache
        edit        edit go.mod from tools or scripts
        graph       print module requirement graph    模块依赖,可以通过graph命令查看某间接依赖是被谁直接依赖的
        init        initialize new module in current directory  
        tidy        add missing and remove unused modules  移除那些没有被使用的模块
        vendor      make vendored copy of dependencies
        verify      verify dependencies have expected content
        why         explain why packages or modules are needed

```

### go mod graph 

### go mod tidy

### go mod init

### go mod vendor

### go mod why  --反正我一直没有使用好这个，不知道为什么


## go get 获取指定模块