package time

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// time.After time.Sleep  time.Tick
func TestFunc(t *testing.T) {
	select {
	case <- time.After(time.Duration(2) * time.Second):
		fmt.Println("time after 2 second done")
	}

	time.Sleep(time.Duration(2) * time.Second)

	// 循环
	tickChan := time.Tick(time.Duration(2) * time.Second)
	for t := range tickChan {
		fmt.Println(t)
	}
	fmt.Println("success")
}

type Interrupt interface {
	Interrupt(ctx context.Context)
}

type After struct {

}
type Sleep struct {

}

// after 可以通过ctx来中断 -- 推荐
func (c *After) Interrupt(ctx context.Context) {
	select {
	case <- time.After(time.Duration(10) * time.Second):
		fmt.Println("time after")
	case <- ctx.Done():
		fmt.Println("ctx done")
	}
	fmt.Println("after done")
}
// sleep 不行
func (c *Sleep) Interrupt(ctx context.Context) {
	time.Sleep(time.Duration(10) * time.Second)
	fmt.Println("sleep done")
}

// compare after and sleep
func TestCompare(t *testing.T) {
	// 当超时时，可以中断的操作
	ctx := context.Background()

	currCtx, cancelFunc := context.WithTimeout(ctx, time.Duration(2)*time.Second)
	defer cancelFunc()
	after := &After{}
	sleep := &Sleep{}

	group :=sync.WaitGroup{}

	group.Add(1)
	go func() {
		after.Interrupt(currCtx)
		defer group.Done()
	}()

	group.Add(1)
	go func() {
		sleep.Interrupt(currCtx)
		defer group.Done()
	}()


	group.Wait()
	fmt.Println("success")
}


