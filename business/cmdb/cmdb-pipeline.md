# cmdb的采集器

采集器: 用于采集三方数据，不同的采集器可以基于不同的方式去采集，比如: rpc、http、https、mysql(表)、等等

不同的采集采集器基于不同的配置,为了可以灵活支持，采用json的方式来实现



## pipeline在代码中的实现逻辑

1、注册了一个pipeline的周期定时任务，每10s执行一次，该周期任务就是用来检测数据库pipeline_config表是否有新增、修改、删除的数据
如果有，那么通过和内存中的数据进行对比，筛选出哪些是新增、哪些是修改、哪些是删除，然后执行对应的安装、卸载安装、卸载操作

2、pipeline_config涉及了整条采集-处理-写入流程
 - 采集相关的配置：表示基于什么方式去哪里采集，采集什么数据
 - 处理相关的配置：表示对采集过来的数据做加工处理  --- 目前没有使用
 - 写入相关的配置：表示将处理后的数据写入到哪里，比如数据库、文件、或者其他的中间件
 同时该配置还涉及到pipeline的定时执行时间、超时时间、是否开启、地域、等
 - 
```text
该任务涉及的相关操作：

新增pipeline: RegisterPipeline

更新pipeline: UnregisterPipeline 再  RegisterPipeline

删除pipeline: UnregisterPipeline
```
### RegisterPipeline
使用了https://github.com/robfig/cron 定时器，类似于Quartz

向"github.com/robfig/cron/v3" 中的cron对象，添加了一个定时任务
```go
// 声明定时器
var corn *cron.Cron
// 创建一个定时器
corn =  cron.New(cron.WithSeconds())

// 开启执行定时任务
corn.Start()

// 添加任务，任务以方法的方式体现，schedule表示任务执行周期 * 3/5 * * *
taskId, err :=corn.AddFunc("schedule", func() {})  

```

### UnregisterPipeline
删除定时任务
```go
corn.Remove(taskId)
```

## pipeline任务的执行
1、将所有的pipeline_config中开启的pipeline进行注册
   - 基于配置生成pipelineContext
   - 将入写入定时任务重corn.addFunc()
2、定时任务周期执行
3、记录定时任务的执行状态
4、执行 pipe.Pipe(ctx)

pipe流程
```go

// 装配 端点: 返回一个接收数据对象next，本质上是一个chan
next, endpointUnit, err := p.assembleEndpoints(ctx, p.Context.EndpointRunners)

// 装配 采集器: 基于数据接收对象装配采集器，将采集的数据写入到chan中
collectorUnit, err := p.assembleCollectors(next, p.Context.CollectorRunners)


var wg sync.WaitGroup  // 等待组，相当于是一个信号量
wg.Add(1) // 添加信号
go func() {
  defer func() {
  wg.Done() // 执行完后释放信号
  log.InfoContextf(ctx, "endpoints run successfully, func quit, taskID: %s", ctx.Value(TaskID))
  }()
  // 执行端点写入
  p.runEndpoints(ctx, endpointUnit)
}()

wg.Add(1)
go func() {
  defer func() {
  wg.Done()
  log.InfoContextf(ctx, "processors run successfully, func quit, taskID: %s", ctx.Value(TaskID))
  }()
  // 执行处理器
  p.runProcessors(ctx, processorUnits)
}()

wg.Add(1)
go func(unit *CollectorUnit) {
  defer func() {
  wg.Done()
  log.InfoContextf(ctx, "collectors run successfully, func quit, taskID: %s", ctx.Value(TaskID))
  }()
  // 执行采集器
  p.runCollectors(ctx, collectorUnit)
}(collectorUnit)

wg.Wait()

```
多个pipe定时执行：

collector(采集数据) --写入--> chan  --读取--> endpoint(写入mysql、文件、其他中间件)

其实在写入到chan后，还有一个加工处理的流程叫做:processor可以对数据进行加工处理，然后再写入到chan中，再被endpoint消费

collector(采集数据) --写入--> chan  --处理--> processor --处理--> chan2 --读取--> endpoint(写入mysql、文件、其他中间件)





