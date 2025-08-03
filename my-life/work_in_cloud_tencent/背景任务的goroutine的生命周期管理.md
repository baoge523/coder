
## 背景任务的goroutine生命周期管理

什么时候需要背景任务：

1、需要定时加载数据库数据 -- 类似于数据预热（起到缓存的效果）
2、异步处理数据，比如chan + select + goroutine 的模式
3、

背景任务的管理：

1、没有很好的被监控，比如背景任务异常时，怎么拿到这个错误
2、背景任务是否优雅退出 -- 其实这里也涉及到goroutine如何优雅退出的问题


在背景任务出错时，能拿到错误信息
背景任务因正常、异常结束时，能够优雅退出
```golang
// 背景任务的接口设计
type BackGroundRunningController interface {
	Start(ctx context) error {}
	
	Stop() error {}
	
	Errors() <- chan error
}

// 使用套路
 a := new_a()
 if err := a.Start(ctx); err != nil {
	 
 }
 defer a.Stop()
 
  b := new_b(a)
if err := b.Start(ctx); err != nil {

}
defer b.Stop()

c := new_c(a)
if err := c.Start(ctx); err != nil {

}
defer c.Stop()


```


### 总结建议
```text
对象职责分离，应明确依赖关系和生命周期，不要全局变量
遵守channel的关闭准则等待完成，而不是草草结束
要和所使用的基础组件相适配，协同完成；
不要各自为政、随意退出进程（局部模块内部log.Fatalf或os.Exit）
```

### 使用例子 - 定时任务
定时任务执行也是背景任务的一种，所以定时任务的框架实现需要满足背景任务的定义

#### 设计步骤
```text
1、定义一个名为Updater的更新器，里面需要包含：
  - 间隔时间，定时任务的间隔时间  -- 使用侧设置
  - 获取数据的行为，为了实现最大的可拓展，所以定义一个返回interface的func  -- 使用侧设置
  - 错误处理，获取数据的行为中，可能会有异常，使用用户侧实现，不可控；为了实现最大方式的可拓展，将其定义为接入ctx、error的func -- 使用侧设置
  - 数据存储, 获取的数据需要存储下来 
  - 取消控制，当遇到意向之外的异常时，需要有能力能个取消任务

2、根据BackGroundRunningController 接口定义，实现对应的方法：Start() Stop() Errors() 等

```
#### 实现方式
```golang
type Interval struct {
	Time         time.Duration // 定时任务的执行时间
	ValueSetter  func(ctx context.Context) (interface{}, error)  // 用户设置，获取数据的func
	ErrorHandler func(context.Context, error)  // 用户设置，当在获取数据过程中出现error时，应该怎么处理

	cancel context.CancelFunc  // 有context取消的能力
	eg     errgroup.Group     // 可收集err的goroutine group，异步执行

	val atomic.Value  // 存储数据  同步操作
}

// todo 实现 start stop 等方法
```

### 使用例子2 -- kafka消费者
消费者就是一个背景任务，该任务消费指定kafka中的数据