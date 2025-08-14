package kafka_sink

import (
	"code_design/data_stream/component"
	"code_design/data_stream/consumer"
	"code_design/data_stream/sink"
	"context"
)

type KafkaSink struct {
	ss     sink.SinkSetting
	config component.ComponentConfig
}

func (ks *KafkaSink) Start(ctx context.Context) error {

	return nil
}

func (ks *KafkaSink) Stop(ctx context.Context) error {

	return nil
}

func (ks *KafkaSink) ConsumerMsg(ctx context.Context, message consumer.Message) error {
	return nil
}
