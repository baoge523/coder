package options

import "testing"

// 使用匿名函数的方式实现option

// 定义一个OptionFunc 该类型是一个接受*Options的函数
type OptionFunc func(*Options)

func WithAddrFunc(addr string) OptionFunc {
	return func(opts *Options) {
		opts.addr = addr
	}
}

func WithHostFunc(host string) OptionFunc {
	return func(opts *Options) {
		opts.host = host
	}
}

func WithPortFunc(port string) OptionFunc {
	return func(opts *Options) {
		opts.port = port
	}
}

func WithDbTypeFunc(dbType string) OptionFunc {
	return func(opts *Options) {
		opts.dbType = dbType
	}
}

func WithCacheFunc(cache bool) OptionFunc {
	return func(opts *Options) {
		opts.cache = cache
	}
}

func WithTimeoutFunc(timeout int64) OptionFunc {
	return func(opts *Options) {
		opts.timeout = timeout
	}
}

func NewDBOptionFunc(optionFunc ...OptionFunc) DB {
	options := &Options{}
	for _, o := range optionFunc {
		o(options)
	}
	return NewDB(options.addr, options.host, options.port, options.dbType, options.cache, options.timeout)
}

func TestOptionFunc(t *testing.T) {
	db := NewDBOptionFunc(
		WithAddrFunc("127.0.0.1:3306"),
		WithHostFunc("127.0.0.1"),
		WithPortFunc("3306"),
		WithDbTypeFunc("mysql"),
		//WithCacheFunc(true),
		//WithTimeoutFunc(20000),
	)
	println(db.string())

}
