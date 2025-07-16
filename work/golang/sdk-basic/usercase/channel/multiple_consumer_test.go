package channel

import (
	"fmt"
	"sync"
	"testing"
)

// TestMultipleConsumer 验证一个channel 多个消费者消费时，是否存在重复消费; 不会存在重复消费
// <- chan 在chan close时，所有的监听能感知到
func TestMultipleConsumer(t *testing.T) {

	c := make(chan int,10)

	wg := sync.WaitGroup{}
	// consumer
	for i := 0; i < 3; i++ {
		wg.Add(1)
		num := i
		go func() {
			defer wg.Done()
			select {
			case data :=  <- c:
				fmt.Printf("current %d go runtine consum data %d \n",num,data)
			}
		}()
	}

	// product
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)

	// wait
	wg.Wait()
	fmt.Println("success")


}
