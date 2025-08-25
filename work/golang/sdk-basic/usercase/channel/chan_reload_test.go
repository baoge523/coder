package channel

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestReload(t *testing.T) {

	ctx, cancelFunc := context.WithCancel(context.Background())

	reload := make(chan int, 10)

	for i := 0; i < 10; i++ {
		reload <- i
	}

	go func() {
		defer cancelFunc()
		<-time.After(10 * time.Second)
		fmt.Println("do cancel ...")
	}()

	for {
		select {
		case <-ctx.Done():
			return // 这里如果使用break 只会跳出select，而不是跳出for
		case <-time.After(2 * time.Second):
			select {
			case <-ctx.Done():
				return
			case <-reload:
				fmt.Println("do reload....")
			}
		}

	}

}
