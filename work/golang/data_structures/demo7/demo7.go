package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var c = make(chan string)

func deal(wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range c {
		fmt.Println(v)
	}
}

func shutdown() {
	close(c)
}

// 目的
// 1.并发处理i，保证每个i只被处理一次
// 2.1秒后系统关闭，优雅的停止服务(不出现panic)

func main() {

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go deal(wg)

	ct, cancel := context.WithCancel(context.Background())
	go func() {
		i := 0
		for {
			select {
			case <-ct.Done():
				return
			default:
				i++
				go func(num int) {
					// 这里

					submit(fmt.Sprintf("%d", num))
				}(i)
			}

		}
	}()

	go func() {
		time.Sleep(time.Second * 1)
		cancel()
		shutdown()
	}()

	//
	wg.Wait()

	select {}
}

func submit(task string) {
	c <- task
}
