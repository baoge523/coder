package memory

type options struct {
	maxWindowSize int
}

func defaultOptions() *options {
	return &options{
		maxWindowSize: 10,
	}
}

// Option is a function that modifies the options
type Option func(*options)

// WithMaxWindowSize sets the maximum window size
func WithMaxWindowSize(maxWindowSize int) Option {
	return func(o *options) {
		o.maxWindowSize = maxWindowSize
	}
}
