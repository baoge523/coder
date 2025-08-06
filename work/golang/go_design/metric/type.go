package metric

import (
	"go.opentelemetry.io/otel/attribute"
)

// Label 标签 key-value 形式，复用的是otel/attribute
type Label = attribute.KeyValue

// LabelKey 标签key 复用的是otel/attribute
type LabelKey = attribute.Key

// Keys 标签键列表，多个标签key，用于统计方式的固定标签key时，使用
type Keys []LabelKey

// GetLabelKeys 返回标签键列表
func (k Keys) GetLabelKeys() []LabelKey {
	return k
}

func (k Keys) asStringsValue(val ...string) []Label {
	if len(val) == 0 {
		return nil
	}
	labelList := make([]Label, 0, len(val))
	for i, currKey := range k {
		// 补全key-value并加入到labellist中
		labelList = append(labelList, currKey.String(val[i]))
	}
	return labelList
}

// LabelsGetter 支持获取标签的对象
type LabelsGetter interface {
	// GetLabels 获取标签对
	GetLabels() []Label
}

// ToLabels 组合得到用于上报的标签列表
func ToLabels(o interface{}, others ...Label) []Label {
	if g, ok := o.(LabelsGetter); ok {
		var ret []Label
		ret = append(ret, g.GetLabels()...)
		ret = append(ret, others...)
		return ret
	}
	return others
}
