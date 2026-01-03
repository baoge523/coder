
## grpc-go 的练习项目
主要是熟悉grpc-go在使用grpc的全流程，从定义pb，生产client-stub、server-stub的整个过程

### 参考官方文档安装插件并知道如何根据pb生成代码


### 了解grpc-go调用过程中，peer和metadata获取的原数据信息

peer 获取调用放的ip信息
```go
// FromContext returns the peer information in ctx if it exists.
func FromContext(ctx context.Context) (p *Peer, ok bool) {
	p, ok = ctx.Value(peerKey{}).(*Peer)
	return
}
```

metadata获取调用信息
```go
// FromIncomingContext returns the incoming metadata in ctx if it exists.
//
// All keys in the returned MD are lowercase.
func FromIncomingContext(ctx context.Context) (MD, bool) {
	md, ok := ctx.Value(mdIncomingKey{}).(MD)
	if !ok {
		return nil, false
	}
	out := make(MD, len(md))
	for k, v := range md {
		// We need to manually convert all keys to lower case, because MD is a
		// map, and there's no guarantee that the MD attached to the context is
		// created using our helper functions.
		key := strings.ToLower(k)
		out[key] = copyOf(v)
	}
	return out, true
}
```

### grpc的timeout时间传递
在golang的context中，timeout本质就是deadline
```golang
// WithTimeout 底层还是使用的deadline
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}
```

但需要注意的是，context是多叉树结构的，每个context都会有自己的deadline，取消方式是： parent context 可以取消 child context
> 例如
> parent context 的deadline 是5秒钟
> child context  的deadline 是30秒钟
> 那么5秒后，parent context被取消后，child context的ctx.Done()也会监听到取消行为
> 

回归正题，grpc-go中context的timeout是如何传递的？  ---- 通过http2中的header grpc-timeout传递的
```go
// client-side =======
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// and use ctx to call


// server-side ======

// 从http2的header中解析grpc-timeout超时时间，计算公式是：超时时间 = client-size的超时时间 - 网络传输时间
case "grpc-timeout":
timeoutSet = true
var err error
timeout, err = decodeTimeout(hf.Value)

// 如果设置了超时时间
if timeoutSet {
s.ctx, s.cancel = context.WithTimeout(ctx, timeout)
} else {
s.ctx, s.cancel = context.WithCancel(ctx)
}

```