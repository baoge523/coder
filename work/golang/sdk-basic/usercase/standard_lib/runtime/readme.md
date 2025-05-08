## 用于管理golang程序运行时状态信息
内存、cpu、goroutine、gc等等

### 环境变量

GOGC
```text
设置什么时候触发gc操作，百分比；
The default is GOGC=100. 
Setting GOGC=off disables the garbage collector entirely
```

GOMEMLIMIT
```text
sets a soft memory limit for the runtime
```

GODEBUG

GOMAXPROCS
```text
指定操作系统的最多线程数量
```

GORACE

GOTRACEBACK
```text
控制错误堆栈的输出数量，当遇到没有捕获的panic或者一个预期之外的运行时错误

 GOTRACEBACK=none   忽略堆栈信息
 GOTRACEBACK=single (the default)  打印当前的goroutine信息，如果没有当前goroutine，比如内容运行时异常，则打印所有的goroutine信息
 GOTRACEBACK=all  所有用户创建的goroutine信息
 GOTRACEBACK=system 和all类似，展示goroutine内部创建信息和栈帧信息
 GOTRACEBACK=crash 和system类似
```

用于go build的环境变量  交叉编译

GOARCH, GOOS, GOPATH, and GOROOT