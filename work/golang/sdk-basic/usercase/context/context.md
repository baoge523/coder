## context的作用
context的作用有如下：
1、传递上下文信息 (链路信息)
2、超时控制
3、cancel 信号的传递，跨goroutine
4、安全凭证 (信息传递)

### cancel 取消信号传递
```text
1、创建基于可以取消的Context
ctx,cancel := context.WithCancel(parentCtx)  
ctx,cancel :=context.WithCancelCause(parentCtx)  // 带有异常的

2、其实是创建的 cancelCtx 对象
通过ctx.Done() 获取 上下文chan，当cancel func被执行时，及对应的chan被close掉，监听上下文chan的routine都会收到信号

3、当前context的cancel被执行时，其所有的子context的ctx.Done()也会收到

```

### 超时控制
```text
1、创建有时间限制的context
// 底层调用 WithDeadline(parent, time.Now().Add(timeout)
ctx, cancelFunc = context.WithTimeout(ctx, time.Second)

2、其实创建的是 timerCtx 继承 cancelCtx

3、当设置的时间触发时，会执行cancel func
```

### 传递上下文信息
```text
1、创建具有存储key-value的context
ctx = context.WithValue(ctx, "hello", "world")

2、其实创建的是 valueCtx

3、查询数据时，递归查询，时间复杂度为O(n)

4、context 是通过组合，形成的树状结构，当当前没有查询到key值时，会去查询其parent context

```