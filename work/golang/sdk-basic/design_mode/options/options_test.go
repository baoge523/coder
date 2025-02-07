package options

import (
	"fmt"
	"testing"
)

// option模式适用于参数比较多的情况下，这样有利于拓展

type DB struct {
	addr    string
	host    string
	port    string
	dbType  string
	cache   bool
	timeout int64
	// ...
}

func NewDB(addr string, host string, port string, dbType string, cache bool, timeout int64) DB {
	db := DB{
		addr:    addr,
		host:    host,
		port:    port,
		dbType:  dbType,
		cache:   cache,
		timeout: timeout,
	}
	return db
}

func NewDBWithOption(option ...Option) DB {
	options := &Options{}
	for _, o := range option {
		o.apply(options)
	}
	return NewDB(options.addr, options.host, options.port, options.dbType, options.cache, options.timeout)
}

func (d DB) string() string {
	return fmt.Sprint(d)
}

func TestNoOption(t *testing.T) {
	db := NewDB("127.0.0.1:3306", "127.0.0.1", "3306", "mysql", true /*timeout*/, 1000)
	println(db.string())
}

// 记录所有的配置项
type Options struct {
	addr    string
	host    string
	port    string
	dbType  string
	cache   bool
	timeout int64
}

// Option 接口，只有一个函数，这样会导致每一个类型都需要实现这个接口
// 可以将 Option 定义成接受*Options的函数
type Option interface {
	apply(*Options)
}

type addrType string

func (a addrType) apply(ops *Options) {
	ops.addr = string(a)
}

func WithAddr(addr string) Option {
	return addrType(addr)
}

type cacheType bool

func (a cacheType) apply(opts *Options) {
	opts.cache = bool(a)
}

func WithCache(cache bool) Option {
	return cacheType(cache)
}

type hostType string

func (a hostType) apply(opts *Options) {
	opts.host = string(a)
}

func WithHost(host string) Option {
	return hostType(host)
}

type portType string

func (a portType) apply(opts *Options) {
	opts.port = string(a)
}

func WithPort(port string) Option {
	return portType(port)
}

type dbType string

func (a dbType) apply(opts *Options) {
	opts.dbType = string(a)
}

func WithDbType(dt string) Option {
	return dbType(dt)
}

// 每一个参数都需要单独定义，是不是太不友好了，是不是可以弄成匿名函数
type timeoutType int64

func (a timeoutType) apply(opts *Options) {
	opts.timeout = int64(a)
}

func WithTimeOut(timeout int64) Option {
	return timeoutType(timeout)
}

func TestUseOption(t *testing.T) {
	db := NewDBWithOption(
		WithAddr("127.0.0.1:3306"),
		WithHost("127.0.0.1"),
		WithPort("3306"),
		WithDbType("mysql"),
		WithCache(true),
		WithTimeOut(1000),
	)
	println(db.string())
}
