package filter_handle

import (
	"fmt"
	"net/http"
	"testing"
)

// 带有context的filter链

type Context struct {
	req   *http.Request
	resp  http.ResponseWriter
	param map[string]string
	next  func()
}

type ctxFilter func(ctx *Context)

func buildChain(filter ...ctxFilter) http.HandlerFunc {
	// chain返回的http.HandlerFunc就是目标的处理器
	return func(writer http.ResponseWriter, request *http.Request) {
		// 构建上下文，并传递下去
		ctx := &Context{
			req:   request,
			resp:  writer,
			param: make(map[string]string),
		}

		// filter chain
		var next func()
		for i := len(filter) - 1; i >= 0; i-- {
			current := filter[i]
			ctx_copy := &Context{
				req:   ctx.req,
				resp:  ctx.resp,
				param: ctx.param,
				next:  ctx.next, // 使用当前构建出来的func
			}
			next = func() {
				current(ctx_copy) // 如果共用一个ctx的next，next会被覆盖，
			}
			ctx.next = next
		}

		if next != nil {
			next()
		}
	}
}

func logCtxFilter(ctx *Context) {
	fmt.Println("logCtxFilter before")
	ctx.param["log"] = "add log param"
	ctx.next()
	fmt.Println("logCtxFilter after")
}

func authCtxFilter(ctx *Context) {
	fmt.Println("authCtxFilter before")
	ctx.param["auth"] = "add auth param"
	ctx.next()
	fmt.Println("authCtxFilter after")
}

func helloWorld(ctx *Context) {
	fmt.Fprintf(ctx.resp, "log = %s; auth = %s", ctx.param["log"], ctx.param["auth"])
}

// 学习到的：
// func 可以递归构建包含，用户的参数如果是局部变量，就是只拷贝；如果是外部变量就是最新的值
func TestCtx(t *testing.T) {
	handler := buildChain(logCtxFilter, authCtxFilter, helloWorld)

	err := http.ListenAndServe("localhost:8081", handler)
	if err != nil {
		fmt.Printf("error %v \n", err)
	}
}
