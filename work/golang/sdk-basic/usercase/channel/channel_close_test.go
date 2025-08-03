package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// TestChannelClose 验证chan被close后，无限的for循环在消费完chan后，会退出吗 -- 会的，这就是优雅退出
func TestChannelClose(t *testing.T) {

	wg := sync.WaitGroup{}
	c := make(chan int, 10)

	wg.Add(1) // 这个需要先加，go routine是异步处理的
	go func() {
		fmt.Println("go start")
		defer wg.Done()
		for v := range c {
			time.Sleep(time.Second)
			fmt.Println(v)
		}
		fmt.Println("go end")
	}()

	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	fmt.Println("close c")

	wg.Wait()
	fmt.Println("success")

}
