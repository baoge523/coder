package errgroup

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

func TestErrGroup(t *testing.T) {

	var eg errgroup.Group
	eg.SetLimit(3) // 限制组内同时存在的最大的活跃协程

	eg.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("sleep 5 run ok")
		return fmt.Errorf("mock error %s", "sleep 2s")
	})

	eg.Go(func() error {
		time.Sleep(time.Second)
		fmt.Println("sleep 1 run ok")
		return fmt.Errorf("mock error %s", "sleep 1s")
	})

	eg.Go(func() error {
		time.Sleep(time.Second)
		fmt.Println("run ok")
		return nil
	})
	eg.Go(func() error {
		time.Sleep(time.Second)
		fmt.Println("1------- run ok")
		return nil
	})
	fmt.Println("block because of limit ")
	eg.Go(func() error {
		time.Sleep(time.Second)
		fmt.Println("2--------run ok")
		return nil
	})
	eg.Go(func() error {
		time.Sleep(time.Second)
		fmt.Println("3--------run ok")
		return nil
	})

	// 如果当前组内活跃的goroutine 和设置的limit一样，那么就不会执行
	// 没有limit则会执行
	tryGo := eg.TryGo(func() error {
		time.Sleep(time.Second)
		fmt.Println("try--------run ok")
		return nil
	})
	fmt.Printf("try go run: %t \n", tryGo)

	// eg.wait 会收集第一个触达的err，但是这里也有一个问题是，它需要等待所有的任务执行完才返回；
	// 其实有时候期望：当存在执行失败了，其他的任务就直接终止吧，不需要再执行了
	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("success")
}

// TestContentErrGroup 当通过context创建errgroup时，errgroup会创建一个withCancelCause，
// 然后在errgroup.Go首次遇到err时执行cancel
// 或者在errgroup.wait 等待所有的goroutine执行完后，执行cancel
func TestContentErrGroup(t *testing.T) {

	eg, ctx := errgroup.WithContext(context.Background())

	eg.Go(func() error {
		select {
		case <-time.After(5 * time.Second):
		case <-ctx.Done():
		}
		fmt.Println("sleep 5 run ok")
		return fmt.Errorf("mock error %s", "sleep 2s")
	})

	eg.Go(func() error {
		time.Sleep(time.Second)
		fmt.Println("sleep 1 run ok")
		return fmt.Errorf("mock error %s", "sleep 1s")
	})

	go func() {
		<-ctx.Done()
		fmt.Println("after context done, what can i should do ~")
	}()

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("success")

	for i := 0; i < 10; i++ {

	}

}
