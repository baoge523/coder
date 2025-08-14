package sink

import (
	"code_design/data_stream/component"
	"context"
)

var _ Factory = (*factory)(nil)

type Sink component.Component

// SinkSetting 设置信息
type SinkSetting struct {
	buildInfo component.BuildInfo
	ID        component.ID

	// other
}

// CreateSink  funcType
type CreateSink func(context.Context, SinkSetting, component.ComponentConfig) (Sink, error)

type Factory interface {
	component.Factory
	CreateSink(ctx context.Context, set SinkSetting, config component.ComponentConfig) (Sink, error)
}

func NewFactory(set SinkSetting, cs CreateSink) Factory {
	return &factory{set: set, cs: cs}
}

type factory struct {
	set SinkSetting
	cs  CreateSink
}

func (f *factory) ID() component.ID {
	return f.set.ID
}
func (f *factory) CreateSink(ctx context.Context, set SinkSetting, config component.ComponentConfig) (Sink, error) {
	return f.cs(ctx, set, config)
}
