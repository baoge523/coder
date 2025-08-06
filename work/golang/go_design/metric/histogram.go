package metric

import "context"

var _ HistogramWithLabels = (*HistogramWithLabelsImpl)(nil)

// Histogram 直方图计数器
type Histogram interface {
	// GetName 获取指标名
	GetName() string
	// GetBuckets 获取分桶设置
	GetBuckets() []float64
	// ObserveWithContext 上报
	ObserveWithContext(context.Context, float64, ...Label)
	// Observer 上报
	Observe(float64, ...Label)

	// NewWithLabels 新建一个标签键固定的计数器
	NewWithLabels(...LabelKey) HistogramWithLabels
}

// HistogramWithLabels 携带指定标签的直方图计数器
type HistogramWithLabels interface {
	// GetName 获取指标名
	GetName() string
	// GetBuckets 获取分桶设置
	GetBuckets() []float64

	// ObserveWithContext 上报
	ObserveWithContext(context.Context, float64, ...string)
	// Observer 上报
	Observe(float64, ...string)
}

// HistogramWithLabelsImpl 携带指定标签的直方图计数器, 具体实现
type HistogramWithLabelsImpl struct {
	Histogram
	Keys
}

// ObserveWithContext 上报
func (h *HistogramWithLabelsImpl) ObserveWithContext(ctx context.Context, v float64, val ...string) {
	h.Histogram.ObserveWithContext(ctx, v, h.Keys.asStringsValue(val...)...)
}

// Observe 上报
func (h *HistogramWithLabelsImpl) Observe(v float64, val ...string) {
	h.ObserveWithContext(context.Background(), v, val...)
}
