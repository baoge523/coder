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
        verify      verify dependencies have expected content  验证依赖是否是期望的内容
        why         explain why packages or modules are needed  

```

### go mod graph 
```bash
# go mod graph 查看mod的依赖图，不支持参数；所以需要通过 | grep 来检查想要查询的版本使用
# 通过 go mod graph | grep golang.org/x/net@v0.22.0 可以查出其是哪个包引入的，同时也能找到其被哪些包使用
go mod graph | grep golang.org/x/net@v0.22.0
```

### go mod tidy
下载 go.mod 文件依赖的包

### go mod init
```bash
go help mod init

# Init initializes and writes a new go.mod file in the current directory
```
在当前目录下执行 go mod init module_name，就会在当前目录下初始化，并生产一个go.mod的文件，go.mod文件中的模块名称为module_name
```text
比如:
cd projects     这里的projects 是目录
go mod init projects  在projects目录下初始化一个模块叫做projects

```


### go mod vendor
在项目中生成vendor目录，该vendor的目的是将项目依赖的包下载到项目中
在执行编译时，可以通过指定使用vendor模式，从而在编译过程中无需联网下载依赖，直接使用本地vendor中的依赖

go mod vendor 常用与无法联网的编译环境

```bash
go mod vendor
go build  -mod=vendor  # 使用vendor模式编译
```

**注意: go mod vendor 无法将依赖的包含cgo相关的(即非golang语言的文件)放入到vendor目录下**
解决方式参考 go-vendor.md 文件


### go mod why  
反正我一直没有使用好这个，不知道为什么


## golang 获取依赖的方式

### go get 

### go install