# 分析golang-sdk 1.21版本中你的net包的实现
官方文档: https://pkg.go.dev/net@go1.21rc2

## client.do方法的流程

### 1、调用链
```text
client.Do(req *Request) -- req请求对象，包含 header、url等信息用于构建请求行、请求头、请求体
  client.send(req *Request, deadline time.Time) -- req 请求对象，deadline请求的超时时间
    send(ireq *Request, rt RoundTripper, deadline time.Time) -- ireq 请求对象，rt 用于发送http请求，通过一个请求获取一个响应， deadline 超时时间
      http.Transport.RoundTrip(req *Request) -- 重要
        Transport.getConn(treq *transportRequest, cm connectMethod) -- treq 包装后的请求对象， cm 管理请求构造key(本来tcp连接缓存)
        http.persistConn.roundTrip(req *transportRequest) -- persistConn 是net.conn的连接对象 req 包装的请求对象
```
以上的调用链中，最重要的方法是http.Transport.RoundTrip，它包含了完整的请求-响应的整体流程

RoundTrip主要分两步：
- 1、获取连接
- 2、通过连接来发送请求

获取连接的步骤包含如下：
- 1、获取连接，通过cm构造的key看是否有空闲连接，如果有，直接返回
- 2、如果没有：当maxConnPreHost <=0表示不限制host的tcp连接数，直接创建连接，否则执行下面的操作
- 3、如果当前创建的连接小于maxConnPreHost那么就直接创建创建，否则就排队等待创建连接
- 4、当新建一个tcp连接的时候，会启动两个线程：readLoop和writeLoop，用于发送和读取连接中的请求和响应


当获取到连接后步骤如下：
- 1、将请求包装后，发送到conn对象中的writeChan中，会被当前连接的writeLoop消费，并发送到server端
- 2、发送请求后，等待响应，readLoop会监听server的响应，并将结果投递到resChan中
- 3、当收到响应结果后，将连接放入到空闲连接池中


### 2、流程图

### 3、关键代码

重要的结构体对象
- client
- Transport
- persistConn

#### ①、client
```go
type Client struct {
	// 用于发送请求得到响应的对象，如果为空，则使用DefaultTransport （重要）
	Transport RoundTripper
	// 重定向
	CheckRedirect func(req *Request, via []*Request) error
	// 用于管理cookie的对象，如果为空，则使用DefaultCookieJar
	Jar CookieJar
	// 超时时间
	Timeout time.Duration
}
```
client对外的相关方法
```text
Do
Get
Head
Post
PostForm
```

使用方式：
```text
func main() {
    client := &http.Client{}
    client.Do()
    
    // 如果想控制TLS、keep-alive、压缩、proxy等其他设置，可以通过Transport来设置
    client2 := &http.Client{
        Transport: &http.Transport{
            MaxIdleConnsPerHost: 100,
        },
    }
    client2.Do()
}
```

#### ②、Transport
transport是RoundTripper的实现者，可以支持http和https请求；
在默认情况下，transport会缓存连接，已达到复用连接的目的；
transport在访问host时，会留下很多打开的连接，这些连接通过CloseIdleConnections、MaxIdleConnsPerHost、DisableKeepAlives来控制；
transport是线程安全的，可以被多个goroutine使用；

为了更好的了解Transport的工作原理，我们可以从如下几个问题开始来分析：
- 1、如何管理连接池
- 2、如何创建连接
- 3、连接在使用完后，是否放回连接池中
- 4、空闲连接什么时机被释放的

DefaultTransport
```go
var DefaultTransport RoundTripper = &Transport{
	Proxy: ProxyFromEnvironment,   // 代理
	DialContext: defaultTransportDialContext(&net.Dialer{  // 创建连接
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}),
	ForceAttemptHTTP2:     true, // 尝试使用http2协议
	MaxIdleConns:          100,  // 最大的空闲连接数
	IdleConnTimeout:       90 * time.Second, // 空闲连接的超时时间
	TLSHandshakeTimeout:   10 * time.Second, // tls握手超时时间
	ExpectContinueTimeout: 1 * time.Second, 
}
```
##### 如何管理连接池
```text
// Transport的结构体中，通过idelConn来保存空闲的连接
idleConn     map[connectMethodKey][]*persistConn


当发送请求时，需要先获取连接对象，在获取连接对象时，首先通过idleConn看是否有空闲的连接，如果有，则直接使用，如果没有，则创建连接；
在发送请求并接受到响应后，会将连接再放回到空闲连接池中
待补充：
检查空闲连接超时的时机有哪些？
1、在获取连接时，如果从空闲连接池中获取到连接，需要检查该连接是否超时
2、是否有定时任务检查？
```


##### 如何创建连接

排队获取空闲连接： <br/>
1、如果关闭了keep-alive，则直接返回 false  <br/>
2、从idleConn中找可用的空闲连接，如果有，先检查连接是否超时，超时了就关闭连接并继续找  <br/>
   如果找到了可用的空闲连接，那么就将连接放到w中，并返回true  <br/>
3、只要没有找到空闲连接，都会将在idleConnWait中等待获取连接  <br/>
```go
func (t *Transport) queueForIdleConn(w *wantConn) (delivered bool) {
	// 如果连接不保活，直接返回 false
	if t.DisableKeepAlives {
		return false
	}

	t.idleMu.Lock()
	defer t.idleMu.Unlock()
	
	t.closeIdle = false

	if w == nil {
		return false
	}

	// 计算连接超时时间的起点
	var oldTime time.Time
	if t.IdleConnTimeout > 0 {
		oldTime = time.Now().Add(-t.IdleConnTimeout)
	}

	// 寻址可用的空闲连接
	if list, ok := t.idleConn[w.key]; ok {
		stop := false
		delivered := false
		for len(list) > 0 && !stop {
			pconn := list[len(list)-1]
			// 如果连接超时了，就释放连接
			tooOld := !oldTime.IsZero() && pconn.idleAt.Round(0).Before(oldTime)
			if tooOld {
				go pconn.closeConnIfStillIdle()
			}
			if pconn.isBroken() || tooOld {
			    // 数组操作
				list = list[:len(list)-1]
				continue
			}
			// 将连接放到w中
			delivered = w.tryDeliver(pconn, nil)
			if delivered {
				if pconn.alt != nil {
					// HTTP/2: multiple clients can share pconn.
					// Leave it in the list.
				} else {
					// HTTP/1: only one client can use pconn.
					// Remove it from the list.
					t.idleLRU.remove(pconn)
					list = list[:len(list)-1]
				}
			}
			stop = true
		}
		if len(list) > 0 {
			t.idleConn[w.key] = list
		} else {
			delete(t.idleConn, w.key)
		}
		if stop {
			return delivered
		}
	}

	// 只要没有获取到空闲连接，都需要放到空闲连接等待的队列中
	if t.idleConnWait == nil {
		t.idleConnWait = make(map[connectMethodKey]wantConnQueue)
	}
	q := t.idleConnWait[w.key]
	q.cleanFront()
	q.pushBack(w)
	t.idleConnWait[w.key] = q
	return false
}
```
重要:idleConnWait何时被执行呢？
```text
触发时机：当连接被当做空闲连接放入到空闲连接池时触发
当连接使用完后，准备将连接放入到空闲连接池时，会先判断是否有请求在idleConnWait中等待连接，如果有，就执行
```

 ----

排队创建连接: <br/>
1、如果没有配置MaxConnsPerHost，则直接创建连接 <br/>
2、如果配置了MaxConnsPerHost，则检查是否超过了限制 <br/>
    如果没有超过限制，则创建连接 <br/>
3、如果超过了限制，则将创建连接的请求放入创建连接的等待队列中 <br/>

```go
func (t *Transport) queueForDial(w *wantConn) {
	w.beforeDial()
	// 如果没有配置MaxConnsPerHost，则直接创建连接
	if t.MaxConnsPerHost <= 0 {
		go t.dialConnFor(w) // 新起一个协程处理创建连接
		return
	}

	t.connsPerHostMu.Lock()
	defer t.connsPerHostMu.Unlock()
    // 如果配置了MaxConnsPerHost，则检查是否超过了限制
    // 如果没有超过限制，则创建连接
	if n := t.connsPerHost[w.key]; n < t.MaxConnsPerHost {
		if t.connsPerHost == nil {
			t.connsPerHost = make(map[connectMethodKey]int)
		}
		t.connsPerHost[w.key] = n + 1
		go t.dialConnFor(w)
		return
	}
    // 如果超过了限制，则将创建连接的请求放入创建连接的等待队列中
	if t.connsPerHostWait == nil {
		t.connsPerHostWait = make(map[connectMethodKey]wantConnQueue)
	}
	q := t.connsPerHostWait[w.key]
	q.cleanFront()
	q.pushBack(w)
	t.connsPerHostWait[w.key] = q
}
```
重要: 创建连接的等待队列何时被执行呢？
```text
触发时机：放存在一个连接被释放的时候
比如：
连接出现异常，关闭连接
空闲连接超时，关闭连接
...
```

创建http连接: <br/>
所以自定义client.transport可以实现指定代理、创建连接的方式等 <br/>
```go
func (t *Transport) dialConn(ctx context.Context, cm connectMethod)
```

```go
func (t *Transport) dial(ctx context.Context, network, addr string) (net.Conn, error) {
	// 优先使用DialContext创建连接
	if t.DialContext != nil {
		c, err := t.DialContext(ctx, network, addr)
		if c == nil && err == nil {
			err = errors.New("net/http: Transport.DialContext hook returned (nil, nil)")
		}
		return c, err
	}
	// 再使用Dial创建连接
	if t.Dial != nil {
		c, err := t.Dial(network, addr)
		if c == nil && err == nil {
			err = errors.New("net/http: Transport.Dial hook returned (nil, nil)")
		}
		return c, err
	}
	return zeroDialer.DialContext(ctx, network, addr)
}
```

#### ③、persistConn

创建persistConn连接对象 <br/>
1、创建net.conn的tcp连接 <br/>
2、判断是否是代理 <br/>
3、启动协程循环读、循环写 <br/>

```go
func (t *Transport) dialConn(ctx context.Context, cm connectMethod) (pconn *persistConn, err error) {
	pconn = &persistConn{
		t:             t,
		cacheKey:      cm.key(),
		reqch:         make(chan requestAndChan, 1),
		writech:       make(chan writeRequest, 1),
		closech:       make(chan struct{}),
		writeErrCh:    make(chan error, 1),
		writeLoopDone: make(chan struct{}),
	}
	
	if cm.scheme() == "https" && t.hasCustomTLSDialer() {

	} else {
		// 创建http连接:net.conn
		conn, err := t.dial(ctx, "tcp", cm.addr())
		if err != nil {
			return nil, wrapErr(err)
		}
		pconn.conn = conn
		if cm.scheme() == "https" {
			var firstTLSHost string
			if firstTLSHost, _, err = net.SplitHostPort(cm.addr()); err != nil {
				return nil, wrapErr(err)
			}
			if err = pconn.addTLS(ctx, firstTLSHost, trace); err != nil {
				return nil, wrapErr(err)
			}
		}
	}

	// Proxy setup. ... 省略代理相关的处理
	
	// 连接读IO
	pconn.br = bufio.NewReaderSize(pconn, t.readBufferSize())
	// 连接写IO
	pconn.bw = bufio.NewWriterSize(persistConnWriter{pconn}, t.writeBufferSize())

	// 循环读
	go pconn.readLoop()
	// 循环写
	go pconn.writeLoop()
	return pconn, nil
}
```

persistConn发送请求：<br/>


### 4、总结
1、在Transport中,idleConnWait、connsPerHostWait的作用是什么？
```text
idleConnWait: 空闲连接的等待队列
connsPerHostWait: 创建连接的等待队列

这两个队列都是用来等待获取连接的

当存在配置了keep-alive=true和MaxIdleConnsPerHost的情况

存在两种情况：
情况1：当前的连接数在0-maxIdleConnsPerHost之间: 
只会使用idleConnWait

情况2：当前的连接数大于maxIdleConnsPerHost:
会使用idleConnWait和connsPerHostWait，当有空闲连接了，就会消费idleConnWait中的请求，当有连接异常、正常关闭了、超时了，就会消费connsPerHostWait中的请求
两者谁先就谁处理，只有一个处理

```
