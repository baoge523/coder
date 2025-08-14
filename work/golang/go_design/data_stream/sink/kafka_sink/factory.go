package kafka_sink

import (
	"code_design/data_stream/component"
	"code_design/data_stream/sink"
	"context"
)

func NewFactory(set sink.SinkSetting) component.Factory {
	return sink.NewFactory(set, CreateSink)
}

func CreateSink(ctx context.Context, ss sink.SinkSetting, config component.ComponentConfig) (sink.Sink, error) {

	return &KafkaSink{ss: ss, config: config}, nil
}
