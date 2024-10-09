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
```linux
go help mod init

Init initializes and writes a new go.mod file in the current directory
```
在当前目录下执行 go mod tidy module_name，就会在当前目录下初始化，并生产一个go.mod的文件，go.mod文件中的模块名称为module_name
```text
比如:
cd projects
go mod init projects  在projects目录下初始化一个模块叫做projects

```


### go mod vendor

### go mod why  --反正我一直没有使用好这个，不知道为什么


## go get 获取指定模块