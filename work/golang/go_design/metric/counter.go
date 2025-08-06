package metric

import "context"

// 该代码的作用是，在编译期间判断CounterWithLabelsImpl是否全实现了CounterWithLabels接口，方便在编译期间排查问题
var _ CounterWithLabels = (*CounterWithLabelsImpl)(nil)

/**
  这里自定义一个接口，用于屏蔽底层的真正的使用者，业务只需要关系接口就好了，同时也是面向接口开发
*/

// Counter 计数器（自增不减）
type Counter interface {
	// GetName 获取指标名
	GetName() string

	// AddWithContext 增加
	AddWithContext(context.Context, int64, ...Label)
	// Add 增加
	Add(int64, ...Label)

	// AddWithContext 具有context上下文的数据加一
	IncWithContext(context.Context, ...Label)
	// Add 数量加一
	Inc(...Label)

	// NewWithLabels 新建一个标签键固定的计数器
	NewWithLabels(...LabelKey) CounterWithLabels
}

// CounterWithLabels 携带指定标签的计数器
type CounterWithLabels interface {
	// GetName 获取指标名
	GetName() string
	// GetLabelKeys 获取对应的指标键
	GetLabelKeys() []LabelKey

	// AddWithContext 新增
	AddWithContext(context.Context, int64, ...string)
	// Add 新增
	Add(int64, ...string)

	// AddWithContext 加1
	IncWithContext(context.Context, ...string)
	// Add 加1
	Inc(...string)
}

// CounterWithLabelsImpl 携带指定标签的计数器, 具体实现
type CounterWithLabelsImpl struct {
	Counter
	Keys
}

// AddWithContext 增加
func (c *CounterWithLabelsImpl) AddWithContext(ctx context.Context, a int64, val ...string) {
	c.Counter.AddWithContext(ctx, a, c.Keys.asStringsValue(val...)...)
}

// Add 增加
func (c *CounterWithLabelsImpl) Add(a int64, val ...string) {
	c.AddWithContext(context.Background(), a, val...)
}

// IncWithContext 加1
func (c *CounterWithLabelsImpl) IncWithContext(ctx context.Context, val ...string) {
	c.Counter.AddWithContext(ctx, 1, c.Keys.asStringsValue(val...)...)
}

// Inc 加1
func (c *CounterWithLabelsImpl) Inc(val ...string) {
	c.AddWithContext(context.Background(), 1, val...)
}
