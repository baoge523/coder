package sync

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestCond(t *testing.T) {

	var lock sync.Mutex // 这个是 *sync.Mutex 实现的 locker接口，所以必须使用指针才是locker的实现者

	cond := sync.NewCond(&lock)

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			num := i
			go func() {
				lock.Lock()
				defer lock.Unlock()
				fmt.Printf("num %d start wait \n",num)
				cond.Wait() // current goroutine
				fmt.Printf("num %d done success \n",num)
			}()
		}
	}
	time.Sleep(time.Second)
	fmt.Println(runtime.NumGoroutine())
	cond.Signal()   // 释放随机一个
	fmt.Println(runtime.NumGoroutine())
	cond.Broadcast()  // 释放所有，类似 close chan
	time.Sleep(time.Second)
	fmt.Println("success done ~~~")
}

// 更推荐使用chan的方式实现cond的效果
func TestChan(t *testing.T) {

	sign := make(chan int,10)

	for i := 0; i < 10; i++ {
		num := i
		go func() {
			fmt.Printf("current %d start ~~ \n",num)
			select {
			case <- sign:
				fmt.Printf("current %d got sign and run \n",num)
			}
		}()
	}
	fmt.Println("no goroutine run ~~~")

	// write sign to sign chan
	sign <- 0  // write one sign to sign chan and random goroutine running
	time.Sleep(time.Second)
	close(sign)  // close the chan and all waiting goroutine running
	time.Sleep(time.Second)
	fmt.Println("main run success ~~~")
}
