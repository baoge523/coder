package consumer

import "context"

type Consumer interface {
	ConsumerMsg(ctx context.Context, message Message) error
}

type Message interface {
	GetMsg() (any, error)
}
