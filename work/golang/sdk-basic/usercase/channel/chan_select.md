## select and chan

### chan
管道，用于处理数据，比如异步处理数据，构建上下游的生产-消费模型处理数据

#### chan的定义

chan type
比如
```golang

var c1 chan int

var c2 chan string

var c3 chan custom_struct

...

```

#### chan的创建方式

Channel: The channel's buffer is initialized with the specified buffer capacity. 
If zero, or the size is omitted, the channel is unbuffered.

```golang

// 基于make函数创建 chan
c1 := make(chan int)      // 没有缓存的chan int，表示存放数据后，如果不消费，再次存放会阻塞
c2 := make(chan int, 0)   // 没有缓存的chan int，表示存放数据后，如果不消费，再次存放会阻塞

c3 := make(chan int,10)   // 缓存大小为10的chan，表示存放10个后，在存放会阻塞，消费完后再消费会阻塞

```

### select
select 一般都是和chan搭配使用，黄金搭档；用来监听一个或者多个chan的读写状态(写入数据、消费数据)，否则会阻塞当前的goroutine，当然如果有default就不会阻塞goroutine


当我们在 Go 语言中使用 select 控制结构时，会遇到两个有趣的现象：
```text
select 能在 Channel 上进行非阻塞的收发操作；
select 在遇到多个 Channel 同时响应时，会随机执行一种情况；
```
