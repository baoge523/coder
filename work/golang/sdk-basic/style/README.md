## golang 编码规范
参考于uber的go编码规范
[uber编码规范](https://github.com/xxjwxc/uber_go_guide_cn?tab=readme-ov-file)

[google_styleguide](https://google.github.io/styleguide/go/decisions)

建议: 所有代码都应该通过golint和go vet的检查并无错误

### interface
永远不要使用指向interface的指针，这个是没有意义的.在go语言中，接口本身就是引用类型，换句话说，接口类型本身就是一个指针

interface合理性验证
>  var _ http.Handler = (*Handler)(nil) // 如果没有实现接口的全部方法，会在编译期报错

接收器 (receiver) 与接口
> 使用值接收器的方法既可以通过值调用，也可以通过指针调用。
> 带指针接收器的方法只能通过指针或 addressable values 调用。
```go
type A interface {
	A1()
	A2()
}
type impl1 struct {
}
func (i *impl1) A1() {
}
func (i impl1) A2() {
}
func main() {
	// 编译会报错
	var a A = impl1{}  // not ok 因为 A1 是指针方法，只能指针访问
	
	// interface 本身就是指针，既可以直接值赋给接口，也可以通过指针赋值给接口
	var b A = &impl1{} // ok 
}
```
### 零值变量
```go
var a int  // 0
var b bool // false
var c string // ""
var d []int // nil
var e map[string]int // nil  读ok，写 panic
var f chan int // nil 读ok,写 panic
var g func(string) int // nil
var h error // nil  error本身是接口，默认是指针类型 为 nil
var x xxx_Struct // 零值对象 
```
#### sync.Mutex 零值对象
```go
var mu sync.Mutex  // 可以直接使用
mu.Lock()
defer mu.Unlock()
```
```go
type SafeCounter struct {
	mu sync.Mutex
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{} // sync.Mutex属性不需要执行，因为零值可用
}
```

### 属性作用域
当需要访问对象中的内部属性时，如：slice、map等，不要直接返回出去，而是返回出去一个副本
```go
type DataManager struct {
	dataA []String
	dataB map[string]string
}

func NewDataManager() *DataManager {
	return &DataManager{
        dataA: make([]string, 0, 100)    // 尽量在使用前初始化，并指定需要的容量，减少在使用中扩容 （可控的扩容）
        dataB: make(map[string]string,100)  // 尽量使用前初始化，减少在使用中扩容 （不可控的扩容，但可以减少）
    }
}
func (dm *DataManager) GetDataA() []string {
	tmp := make([]string,len(dm.dataA))
	copy(tmp,dm.dataA)
	return tmp // 返回副本，避免直接返回对象，导致对象被修改
}

func (dm *DataManager) GetDataB() map[string]string {
	tmp := make(map[string]string,len(dm.dataB))
	for k,v := range dm.dataB {
		tmp[k] = v
    }
	return tmp // 返回副本，避免直接返回对象，导致对象被修改
}

```

### 使用defer释放资源
1、锁资源
2、io 的close方法
3、请求后得到的响应body流
4、tcp连接
5、数据库连接
6、文件句柄
7、网络连接
8、redis连接

### channel 的最佳使用
channel 通常 size 应为 1 或是无缓冲的。默认情况下，channel 是无缓冲的，其 size 为零。
任何其他尺寸都必须经过严格的审查

- channel 一般用于协程之间的通信
- channel 通常也会搭配select使用，select 可以同时监听多个 channel 的数据，当有数据可读时，就读取数据，并处理。
- 当需要当代一个go routine 执行完成时，可以使用channel 来通知；等待多个goroutine 执行完成时，可以使用sync.WaitGroup 


### 所有有关时间的处理，都需要使用time来操作
时间处理本身就一个十分复杂的事情：地区、时区、等等


### 错误类型 errors
尽可能的通过errors.Is 或 errors.As 来判断错误类型 <br/>
如果是静态的错误信息，可以使用errors.new来创建一个error；如果是动态的，可以通过fmt.Errorf来创建一个error <br/>
使用 fmt.Errorf 搭配 %w 将错误添加进上下文后返回 <br/>
在使用 fmt.Errorf 时，尽可能将当前的信息携带到错误中，以便构成错误链、方便错误排查 <br/>


### 断言处理
类型断言 将会在检测到不正确的类型时，以单一返回值形式返回 panic。 因此，请始终使用“逗号 ok”习语。 <br/>
```go
t, ok := i.(string)
if !ok {
  // 优雅地处理错误
}
```


### 不要一劳永逸地使用 goroutine
Goroutines 是轻量级的，但它们不是免费的： 至少，它们会为堆栈和 CPU 的调度消耗内存。 
虽然这些成本对于 Goroutines 的使用来说很小，但当它们在没有受控生命周期的情况下大量生成时会导致严重的性能问题。 
具有非托管生命周期的 Goroutines 也可能导致其他问题，例如防止未使用的对象被垃圾回收并保留不再使用的资源。

一般来说，每个 goroutine:
- 必须有一个可预测的停止运行时间； 或者
- 必须有一种方法可以向 goroutine 发出信号它应该停止
```go
var (
  stop = make(chan struct{}) // 告诉 goroutine 停止
  done = make(chan struct{}) // 告诉我们 goroutine 退出了
)
go func() {
  defer close(done)
  ticker := time.NewTicker(delay)
  defer ticker.Stop()
  for {
    select {
    case <-tick.C:
      flush()
    case <-stop:
      return
    }
  }
}()
// 其它...
close(stop)  // 指示 goroutine 停止
<-done       // and wait for it to exit
```