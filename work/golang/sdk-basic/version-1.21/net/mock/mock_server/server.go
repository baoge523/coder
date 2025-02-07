package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	serverTest()
}


func serverTest()  {

	http.HandleFunc("/abcd", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("----------- request start-----------")
		body := r.Body

		for s, strings := range r.Header {
			fmt.Println(s, strings)
		}

		contentLenStr := r.Header.Get("Content-Length")
		contentLength, err := strconv.Atoi(contentLenStr)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		reader := bufio.NewReader(body)
		bufBody := make([]byte,contentLength)
		reader.Read(bufBody)
		defer body.Close()

		fmt.Println("request body:", string(bufBody))

		w.Write([]byte("ok"))
		fmt.Println("----------- request end -----------")
	})

	fmt.Println("----------- server start -----------")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	fmt.Println("----------- server end -----------")
}