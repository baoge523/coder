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

### chan 的初始化与使用
1、未初始化的chan，读取数据，会阻塞
```golang
var c1 chan int
<-c1  // 阻塞
```
2、未初始化的chan，使用select case + default时，会执行default
```text
var c1 chan int 
select {
  case <- c1:
  default:
     fmt.println("do")
}
```
3、对未初始化或者已经close了的chan进行写操作，会panic

4、可以对已经close了的chan(有无缓存都可以)进行读操作

5、chan的关闭原则
```text
确认所有的写入方都不再写入的时候，就可以关闭chan了

但怎么确认所有的写入方都不再写入了呢？ -- 因为在很多情况下都是多个写入方

解决方式：
  可以基于明确的组合方式的行为来确认，比如通过defer
 
  a := new_A()
  defer a.close()
  
  b := new_B(a)  // b 是生产者，往a中的chan生产数据
  defer b.close()
  
  c := new_C(a)  // c 也是生产者，往a中的chan生产数据
  defer c.close()

因为defer是栈(FILO)的方式执行，所以在执行到 defer a.close()时，b、c已经close掉了，即表面b c不再向a中的chan生产数据
所以在a.close()中，就可以直接关闭a中的chan；当然如果此时a中的chan还有数据的话，可以继续消费
```

### chan struct{}
struct{}对象不会占用内容，所以在需要通过chan做信号传递时，可以声明使用 chan struct{}
```text
c2 := make(chan struct{})
```