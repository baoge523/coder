## open-telemetry

trace的参考文档
https://opentelemetry.io/docs/languages/go/getting-started/

### 入门使用
1、安装
```linux
go get "go.opentelemetry.io/otel" \
  "go.opentelemetry.io/otel/exporters/stdout/stdoutmetric" \
  "go.opentelemetry.io/otel/exporters/stdout/stdouttrace" \
  "go.opentelemetry.io/otel/exporters/stdout/stdoutlog" \
  "go.opentelemetry.io/otel/sdk/log" \
  "go.opentelemetry.io/otel/log/global" \
  "go.opentelemetry.io/otel/propagation" \
  "go.opentelemetry.io/otel/sdk/metric" \
  "go.opentelemetry.io/otel/sdk/resource" \
  "go.opentelemetry.io/otel/sdk/trace" \
  "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"\
  "go.opentelemetry.io/contrib/bridges/otelslog"
```

2、构建一个全局的Tracer Provider   --- 注意：这里根据需求来，也可以初始化其他的，比如log、metric(只不过我这里不需要)

这里需要指定trace数据写入到哪里，比如：控制台、文件、外部存储系统等
```go
// setupOTelSDK bootstraps the OpenTelemetry pipeline.
// If it does not return an error, make sure to call shutdown for proper cleanup.
func setupOTelSDK(ctx context.Context) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error

	// shutdown calls cleanup functions registered via shutdownFuncs.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	// handleErr calls shutdown for cleanup and makes sure that all errors are returned.
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}
	// Set up trace provider.
	tracerProvider, err := newTracerProvider()
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)
	return
}

func newTracerProvider() (*trace.TracerProvider, error) {
	traceExporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter,
			// Default is 5s. Set to 1s for demonstrative purposes.
			trace.WithBatchTimeout(time.Second)),
	)
	return tracerProvider, nil
}

```

3、构建http服务 -- 注意，这里需要通过otel的handle包装一下
```go

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() (err error) {
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry. 设置全局的 tracer provider
	otelShutdown, err := setupOTelSDK(ctx)
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	// Start HTTP server.
	srv := &http.Server{
		Addr:         ":8080",
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      newHTTPHandler(),  // 通过otel包装后的 handle 返回给http监听使用
	}
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err = <-srvErr:
		// Error when starting HTTP server.
		return
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	err = srv.Shutdown(context.Background())
	return
}



func newHTTPHandler() http.Handler {
	mux := http.NewServeMux()

    // 包装一下
	// handleFunc is a replacement for mux.HandleFunc
	// which enriches the handler's HTTP instrumentation with the pattern as the http.route.
	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		// Configure the "http.route" for the HTTP instrumentation.
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		mux.Handle(pattern, handler)
	}

	// Register handlers.
	handleFunc("/rolldice/", rolldice)
	handleFunc("/rolldice/{player}", rolldice)

	// Add HTTP instrumentation for the whole server.
	handler := otelhttp.NewHandler(mux, "/")
	return handler
}
```

4、构建handle处理器
```go
const name = "go.opentelemetry.io/otel/example/dice"

var (
	tracer = otel.Tracer(name)
)

func rolldice(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "roll")  // 这里可以通过options来设置更多的初始化attr信息
	defer span.End()

	roll := 1 + rand.Intn(6)

	var msg string
	if player := r.PathValue("player"); player != "" {
		msg = fmt.Sprintf("%s is rolling the dice", player)
	} else {
		msg = "Anonymous player is rolling the dice"
	}
	logger.InfoContext(ctx, msg, "result", roll)

	rollValueAttr := attribute.Int("roll.value", roll)
	span.SetAttributes(rollValueAttr)  // 设置属性
	rollCnt.Add(ctx, 1, metric.WithAttributes(rollValueAttr))

	resp := strconv.Itoa(roll) + "\n"
	if _, err := io.WriteString(w, resp); err != nil {
		log.Printf("Write failed: %v\n", err)
	}
}
```

### 生产应用级别 -- 如何初始化
```go

import (
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
)


o := defaultSetupOptions()     // 全部的选项信息
	for _, opt := range options {  // 参数option配置
		opt(o)
	}
	
	var opts []sdktrace.TracerProviderOption
	opts = append(opts, sdktrace.WithSampler(o.sampler))
	opts = append(opts, sdktrace.WithSpanProcessor(
		trace.NewDeferredSampleProcessor(
			trace.NewBatchSpanProcessor(exp, o.batchSpanOption...), o.deferredSampler)))
	opts = append(opts, sdktrace.WithSpanProcessor(zpage.GetZPageProcessor()))

	opts = append(opts, sdktrace.WithResource(res))
	opts = append(opts, sdktrace.WithIDGenerator(o.idGenerator))

	traceProvider := sdktrace.NewTracerProvider(opts...)  // 根据
	otel.SetTracerProvider(traceProvider)

    otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{},
    propagation.Baggage{}))
```


### open-telemetry使用中遇到的问题

#### 问题一：发现trance id 全是0 
这是一个传播配置或代码集成的问题


可能的问题1：  你的代码初始化了 Tracer Provider，但却没有调用 otel.SetTextMapPropagator()

发生过程：

1.一个请求携带了正确的 traceparent头到达你的服务。

2.你的 HTTP 处理器（如 otelhttp.Handler）试图从请求头中提取上下文。

3.但由于没有设置全局传播器，SDK 不知道如何解析 traceparent头，提取失败。

4.提取失败后，处理器会认为这是一个全新的请求，于是尝试创建一个新的 Trace。

5.然而，在创建过程中也可能因为上下文异常，最终生成了一个全零的无效 Trace ID。

解决方式： 
```go
import "go.opentelemetry.io/otel/propagation"

func main() {
    tp := initTracer() // 你的初始化函数
    // 关键：设置传播器！！！
    propagator := propagation.NewCompositeTextMapPropagator(
        propagation.TraceContext{},
        propagation.Baggage{},
    )
    otel.SetTextMapPropagator(propagator)
    // ... 其余代码
}
```

可能的问题2：在中间件(filter)例如日志、认证、限流 等这些中间件中，中断了上下文导致的


解决方案：
> 确保 otelhttp.NewHandler包装的中间件是第一个（或尽可能早）执行的中间件

可能的问题3：自己手动的创建了一个新的context，没有继承之前的context导致


### open-telemetry中的 trace id的传递方式
OpenTelemetry 在 HTTP 和 RPC 场景下，都是通过协议的 Header（或元数据）来传递 Trace ID 的。 这是实现分布式上下文传播的标准方法。