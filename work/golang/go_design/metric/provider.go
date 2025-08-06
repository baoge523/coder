package metric

import (
	"context"
	"sync"
)

/*
*

		该文件用于向使用侧提供一个创建不同统计方式的入口
		提供一个内置的provider接口，同时也提供了各种创建指标统计方式的方法：比如创建counter、gauge、histogram 等
	    内置了一个默认实现的provider,这个内置的provider提供了空实现
	    提供了一个注册provider func,可以通过该func注册一个用户想要的provider,该provider是具体的counter、gauge、histogram的提供者
*/
var (
	dp = &delegationProxy{provider: &noonProvider{}}
)

// ----------------配置 -----------------------------

// NewOptions 新建统计对象时候的参数
type NewOptions struct {
	Description string    // 计数器的描述信息
	Buckets     []float64 // 分桶设置，仅直方图场景下有效
}

// Apply 执行修改函数
func (o *NewOptions) Apply(f ...NewOption) {
	for _, currFun := range f {
		currFun(o)
	}
}

// NewOption 设置参数
type NewOption func(*NewOptions)

// WithDescription 指定描述
func WithDescription(d string) NewOption {
	return func(no *NewOptions) {
		no.Description = d
	}
}

// WithBuckets 设置分桶
func WithBuckets(b []float64) NewOption {
	return func(no *NewOptions) {
		no.Buckets = b
	}
}

// ----------------配置 -----------------------------

// -----------------新增指标统计的提供者(需要在组装时提供具体实现)----------------------------

// Provider 提供者
type Provider interface {
	NewCounter(string, ...NewOption) Counter
	NewGauge(string, ...NewOption) Gauge
	NewHistogram(string, ...NewOption) Histogram
}

// RegisterGlobalProvider  提供一个func能够注入真正干活的provider;通过这些暴露的方法可以执行创建
func RegisterGlobalProvider(p Provider) {
	dp.provider = p
}

// GetGlobalProvider 获取全局提供者
func GetGlobalProvider() Provider {
	return dp
}

// ------------- 暴露给外部直接使用的 ----------------

// NewCounter 新建只增不减的计数器
func NewCounter(n string, o ...NewOption) Counter {
	return dp.NewCounter(n, o...)
}

// NewGauge 新建可增可减的统计对象
func NewGauge(n string, o ...NewOption) Gauge {
	return dp.NewGauge(n, o...)
}

// NewNewHistogram 新建直方图统计对象
func NewNewHistogram(n string, o ...NewOption) Histogram {
	return dp.NewHistogram(n, o...)
}

// ------------------ 内部代理器，主要就是代理真正干活的提供器 -------------------

// provider 的代理器，主要用于代理真实的提供器(真正干活的provider需要在使用前，组装)
type delegationProxy struct {
	provider Provider
}

func (dp *delegationProxy) NewCounter(n string, o ...NewOption) Counter {
	return &delegationCounter{name: n, opts: o}
}
func (dp *delegationProxy) NewHistogram(n string, o ...NewOption) Histogram {
	return &delegationHistogram{name: n, opts: o}
}
func (dp *delegationProxy) NewGauge(n string, o ...NewOption) Gauge {
	return &delegationGauge{name: n, opts: o}
}

// --------------------- 下面就是各个代理器的真实实现(本质还是代理)-------------------------

// 下面包含:noonprovider,delegationprovider,noonCounter,delegationCounter,noonGauge,delegationGauge,noonHistogram,delegationHistogram

type delegationCounter struct {
	ori Counter // 真正实现的counter功能的代理；怎么来的，通过真正的provider创建而来的

	name string
	opts []NewOption
	once sync.Once
}

func (c *delegationCounter) init() {
	c.once.Do(func() { // 只初始化一次，
		c.ori = dp.provider.NewCounter(c.name, c.opts...)
	})
}

func (c *delegationCounter) GetName() string {
	return c.name
}

func (c *delegationCounter) AddWithContext(ctx context.Context, v int64, labels ...Label) {
	c.init()
	c.ori.AddWithContext(ctx, v, labels...)
}

func (c *delegationCounter) Add(v int64, labels ...Label) {
	c.init()
	c.ori.Add(v, labels...)
}

func (c *delegationCounter) IncWithContext(ctx context.Context, labels ...Label) {
	c.init()
	c.ori.IncWithContext(ctx, labels...)
}

func (c *delegationCounter) Inc(labels ...Label) {
	c.init()
	c.ori.Inc(labels...)
}

func (c *delegationCounter) NewWithLabels(k ...LabelKey) CounterWithLabels {
	return &CounterWithLabelsImpl{Counter: c, Keys: k}
}

type delegationGauge struct {
	ori Gauge

	name string
	opts []NewOption
	once sync.Once
}

func (g *delegationGauge) init() {
	g.once.Do(func() {
		g.ori = dp.provider.NewGauge(g.name, g.opts...)
	})
}

func (g *delegationGauge) GetName() string {
	return g.name
}

func (g *delegationGauge) SetWithContext(ctx context.Context, v float64, labels ...Label) {
	g.init()
	g.ori.SetWithContext(ctx, v, labels...)
}

func (g *delegationGauge) Set(v float64, labels ...Label) {
	g.init()
	g.ori.Set(v, labels...)
}

func (g *delegationGauge) NewWithLabels(k ...LabelKey) GaugeWithLabels {
	return &GaugeWithLabelsImpl{Gauge: g, Keys: k}
}

type delegationHistogram struct {
	ori Histogram

	name string
	opts []NewOption
	once sync.Once
}

func (h *delegationHistogram) init() {
	h.once.Do(func() {
		h.ori = dp.provider.NewHistogram(h.name, h.opts...)
	})
}

func (h *delegationHistogram) GetName() string {
	return h.name
}

func (h *delegationHistogram) GetBuckets() []float64 {
	h.init()
	return h.ori.GetBuckets()
}

func (h *delegationHistogram) ObserveWithContext(ctx context.Context, v float64, labels ...Label) {
	h.init()
	h.ori.ObserveWithContext(ctx, v, labels...)
}

func (h *delegationHistogram) Observe(v float64, labels ...Label) {
	h.init()
	h.ori.Observe(v, labels...)
}

func (h *delegationHistogram) NewWithLabels(k ...LabelKey) HistogramWithLabels {
	return &HistogramWithLabelsImpl{Histogram: h, Keys: k}
}

type noonProvider struct{}

func (noonProvider) NewCounter(n string, o ...NewOption) Counter {
	return &noopCounter{name: n}
}

type noopCounter struct {
	name string
}

func (c *noopCounter) GetName() string {
	return c.name
}

func (c *noopCounter) AddWithContext(context.Context, int64, ...Label) {}

func (c *noopCounter) Add(int64, ...Label) {}

func (c *noopCounter) IncWithContext(context.Context, ...Label) {}

func (c *noopCounter) Inc(...Label) {}

func (c *noopCounter) NewWithLabels(k ...LabelKey) CounterWithLabels {
	return &CounterWithLabelsImpl{Counter: c, Keys: k}
}

type noopGauge struct {
	name string
}

func (noonProvider) NewGauge(n string, o ...NewOption) Gauge {
	return &noopGauge{name: n}
}

func (g *noopGauge) GetName() string {
	return g.name
}

func (g *noopGauge) SetWithContext(context.Context, float64, ...Label) {}

func (g *noopGauge) Set(float64, ...Label) {}

func (g *noopGauge) NewWithLabels(k ...LabelKey) GaugeWithLabels {
	return &GaugeWithLabelsImpl{Gauge: g, Keys: k}
}

type noopHistogram struct {
	name string
	b    []float64
}

func (noonProvider) NewHistogram(n string, o ...NewOption) Histogram {
	opts := &NewOptions{}
	opts.Apply(o...)
	return &noopHistogram{name: n, b: opts.Buckets}
}

func (h *noopHistogram) GetName() string {
	return h.name
}

func (h *noopHistogram) GetBuckets() []float64 {
	return h.b
}

func (h *noopHistogram) ObserveWithContext(context.Context, float64, ...Label) {}

func (h *noopHistogram) Observe(float64, ...Label) {}

func (h *noopHistogram) NewWithLabels(k ...LabelKey) HistogramWithLabels {
	return &HistogramWithLabelsImpl{Histogram: h, Keys: k}
}
