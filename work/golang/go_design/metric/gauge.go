package metric

import "context"

var _ GaugeWithLabels = (*GaugeWithLabelsImpl)(nil)

// Gauge 可增可减的统计对象
type Gauge interface {
	// GetName 获取指标名
	GetName() string

	// SetWithContext 设置
	SetWithContext(context.Context, float64, ...Label)
	// Set 设置
	Set(float64, ...Label)

	// NewWithLabels 新建一个标签键固定的计数器
	NewWithLabels(...LabelKey) GaugeWithLabels
}

// GaugeWithLabels 携带指定标签的可增可减的统计对象
type GaugeWithLabels interface {
	// GetName 获取指标名
	GetName() string
	// GetLabelKeys 获取对应的指标键
	GetLabelKeys() []LabelKey

	// SetWithContext 设置
	SetWithContext(context.Context, float64, ...string)
	// Set 设置
	Set(float64, ...string)
}

// GaugeWithLabelsImpl 携带指定标签的可增可减的统计对象, 具体实现
type GaugeWithLabelsImpl struct {
	Gauge
	Keys
}

// SetWithContext 设置
func (g *GaugeWithLabelsImpl) SetWithContext(ctx context.Context, v float64, val ...string) {
	g.Gauge.SetWithContext(ctx, v, g.Keys.asStringsValue(val...)...)
}

// Set 设置
func (g *GaugeWithLabelsImpl) Set(v float64, val ...string) {
	g.SetWithContext(context.Background(), v, val...)
}
