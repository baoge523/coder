package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{7, 8, 9, 3, 4, 5}

	lenA := len(a) - 1
	lenB := len(b) - 1

	total := lenA + lenB

	i := 0
	j := 0
	target := -1
	for i <= total {
		var valA, valB int
		if i <= lenA {
			valA = a[i]
		} else {
			valA = b[i-lenA]
		}

		if j <= lenB {
			valB = b[j]
		} else {
			valB = a[j-lenB]
		}
		i++
		j++
		if valA == valB {
			target = valA
			break
		}

	}
	fmt.Println(target)
}

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

func main1() {

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
