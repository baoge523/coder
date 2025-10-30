package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	test()

}

func test() {
	wg := sync.WaitGroup{}
	num1 := 0
	target := 5
	for i := 0; i < 10; i++ {
		wg.Add(1)
		temp := num1
		go func() { // 一般启动一个goroutine都会有一定的时间（等待cpu的调度），所以当执行是num1早就被改成10了，所以拿到的数据就是10
			defer wg.Done()

			if temp > target {
				fmt.Println("done")
				return
			}
		}()
		num1++
	}
	wg.Wait()
	fmt.Println("ok")
}

func test1() {
	a := []byte("aaa/bbbb")
	index := bytes.IndexByte(a, byte('/'))
	b := a[:index]
	c := a[index+1:]

	b = append(b, "ccc"...)

	// 数组通过[:]分割出来的数组，会共用内存
	fmt.Println(string(a)) // aaacccbb
	fmt.Println(string(b)) // aaaccc
	fmt.Println(string(c)) // ccbb
}
