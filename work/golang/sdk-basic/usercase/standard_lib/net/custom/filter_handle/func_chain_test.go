package filter_handle

import (
	"fmt"
	"net/http"
	"testing"
)

// writer a filter chain

// filter 是对http.HandlerFunc的包装
type filter func(handler http.HandlerFunc) http.HandlerFunc

// FilterChain 本身就是一个filter，包装器模式，一个filter包一个filter，最后包装具体的handler；类似于洋葱
// 相当于 FilterChain 最终返回出去的是一个filter链，这个filter链需要包装最终处理的handler
func FilterChain(f ...filter) filter {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		for i := len(f) -1; i >=0 ; i-- {
			handler = f[i](handler)
		}
		return handler
	}
}


func businessHandle(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("businessHandle do  request url = %s \n",request.URL.RequestURI())
}

func loggerFilter (handler http.HandlerFunc) http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {

		fmt.Println("logger before handle")

		handler(writer,request)

		fmt.Println("logger after handler")
	}
}

func authFilter (handler http.HandlerFunc) http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {

		fmt.Println("auth before handle")

		handler(writer,request)

		fmt.Println("auth after handler")
	}
}

func TestChain(t *testing.T) {
	filterChain := FilterChain(loggerFilter, authFilter)

	err := http.ListenAndServe("localhost:8080", filterChain(businessHandle))
	if err != nil {
		fmt.Println("error")
	}
}

// 创建一个别名，是一个新类型，可以有自己的方法
type a []string

func (f a) toString() string {
	return fmt.Sprintf("%+v",f)
}

// 只是创建一个别名，不是新类型，不能有自己的方法
type b = []string


