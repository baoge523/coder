package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestNoBufChannel(t *testing.T) {
	c1 := make(chan int)
	go func() {
		fmt.Println("go func write 1 to c1")
		c1 <- 1  // 没有缓存，会阻塞
		fmt.Println("go func already write 2 to c1")
	}()

	time.Sleep(time.Second * 5)
	select {
    case v := <-c1:
        fmt.Println("main read v from c1", v)
    default:
		fmt.Println("default")
    }
}

func TestOneBufChannel(t *testing.T) {
	c1 := make(chan int, 1)

    go func() {
        fmt.Println("go func write 1 to c1")
        c1 <- 1  // 有缓存，不会阻塞
        fmt.Println("go func already write 2 to c1")
    }()

    time.Sleep(time.Second * 5)
    select {
    case v := <-c1:
        fmt.Println("main read v from c1", v)
    default:
        fmt.Println("default")
    }
}

// 等待单线程执行完可以使用channel
// 等待多线程执行完可以使用sync.WaitGroup
func TestCloseChannel(t *testing.T) {
	// 无缓存channel
	c1 := make(chan int)
	go func() {
		defer close(c1) // 关闭c1
		fmt.Println("go func running")
		time.Sleep(time.Second * 5)
	}()
	<- c1
	fmt.Println("success")
}

func TestNilChannel(t *testing.T) {
    var c1 chan int
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("sleep ok")
		c1 = make(chan int,1)
		c1 <- 1
	}()
	fmt.Println("main running")
    <- c1 // 需要初始化之后使用，不然会出现死锁
	fmt.Println("main running ok")
}