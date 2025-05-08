package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {

	group := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		num := i
		group.Add(1)
		go func() {
			// group.Add(1) 不能保证 current exec before main
			fmt.Printf("current goroutine %d start \n",num)
			defer group.Done()
			fmt.Printf("current goroutine %d doing \n",num)
		}()
	}

	fmt.Printf("main wait for other goroutine finish \n")
	group.Wait()
	fmt.Printf("main success \n")

}
