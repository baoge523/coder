package supervisor_agent

import (
	"go_ai/demo/a2a/agents/supervisor_agent/memory"
)

type options struct {
	memoryOptions []memory.Option
}

func defaultOptions() *options {
	return &options{}
}

// Option is a function that modifies the options
type Option func(*options)

func WithMemoryOptions(memoryOpts ...memory.Option) Option {
	return func(o *options) {
		o.memoryOptions = memoryOpts
	}
}
