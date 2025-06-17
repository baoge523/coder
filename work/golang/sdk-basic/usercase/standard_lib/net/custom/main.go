package main

import (
	"fmt"
	"net/http"
)

// 上下文结构体
type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Params  map[string]string
	Next    func()
}

// Filter类型定义
type Filter func(*Context)

// 链式构建函数
func BuildFilterChain(filters ...Filter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 构建调用链
		ctx := &Context{
			Writer:  w,
			Request: r,
			Params:  make(map[string]string),
		}
		var next func()
		for i := len(filters) - 1; i >= 0; i-- {
			current := filters[i]
			newCtx := &Context{
				Writer:  ctx.Writer,
				Request: ctx.Request,
				Params:  ctx.Params, // 透传params
				Next:    next,       // 如果只是使用外层ctx作为参数，会导致调用时所有的next都是相同的，然后会造成栈溢出
			}
			next = func() { current(newCtx) }
		}

		if next != nil {
			next()
		}
	}
}

// 日志过滤器
func Logging(ctx *Context) {
	fmt.Printf("Start request: %s %s\n", ctx.Request.Method, ctx.Request.URL.Path)
	ctx.Next()
	fmt.Println("End request")
}

// 参数解析过滤器
func ParseParams(ctx *Context) {
	// 模拟解析URL参数
	ctx.Params["id"] = "123"
	ctx.Next()
}

// 业务处理过滤器
func HelloHandler(ctx *Context) {
	id := ctx.Params["id"]
	fmt.Fprintf(ctx.Writer, "Hello, ID: %s", id)
}

func main() {
	// 构建过滤器链
	handler := BuildFilterChain(Logging, ParseParams, HelloHandler)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
