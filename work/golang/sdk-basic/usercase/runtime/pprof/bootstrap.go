package main

import (
	"fmt"
	"log"
	"net/http"
	 _ "net/http/pprof"
)

// https://pkg.go.dev/net/http/pprof@go1.20
func main() {

	// http://localhost:6060/debug/pprof/
	go func() {
		// 使用net/http/pprof
		fmt.Println("start ")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	mockUseMemory := make([]byte,1024 *1024)

	if len(mockUseMemory) == 100 {
		fmt.Println("mockUseMemory len = 100")
	}

	// 模拟业务
	_ = http.ListenAndServe("localhost:8080", &HelloWorld{})
	fmt.Println("over")
}

type HelloWorld struct {

}

func (h *HelloWorld) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Printf("url = %s",req.RequestURI)

	_, err := resp.Write([]byte("success"))
	if err != nil {
		fmt.Printf("error = %v",err)
	}
	return
}