package context

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCancel(t *testing.T) {

	// 这里的ctx 是 cancelCtx
	ctx, cancelFunc := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "hello", "world") // ctx 是valueCtx，其中有一个context用来存放parent，key存放入参key，value存放入参value
	ctx = context.WithValue(ctx, "key1", "value2")
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		currentNum := i
		go func() {
			defer wg.Done()
			fmt.Printf("currentNum = %d  runing \n", currentNum)
			select {
			case <-ctx.Done(): // 这里的ctx.Done(),本质是调用的cancelCtx.Done(),是返回的cancelCtx.done字段存放的值
				fmt.Printf("currentNum = %d stop run\n", currentNum)
			}
		}()
	}
	time.Sleep(time.Second)
	cancelFunc() // 这里执行时，会执行会将cancelCtx.done里面的chan进行close,这样所有监听的<-ctx.Done() 都会收到消息
	wg.Wait()
	err := ctx.Err()
	if err != nil {
		fmt.Println(err.Error())
	}
	value := ctx.Value("hello") // 判断context类型，然后递归查询，O(n)的时间复杂度
	fmt.Println(value)
	value2 := ctx.Value("key1")
	fmt.Println(value2)
	fmt.Println("success~~~")

}

// TestMultipleContext 得出结论：
// 1、cancelCtx 当是parent执行cancel时，parent and children 其对象的Done() chan都会被close
// 2、cancelCtx 当是child执行cancel时，只有该child的Done() chan会被close
func TestMultipleContext(t *testing.T) {

	// backgroundCtx 继承(golang中是组合) emptyCtx； emptyCtx实现了context.Context接口的四个方法(Deadline|Done|Err|Value)但都是空实现
	background := context.Background()

	// ctxParent 是 cancelCtx
	// cancelFunc 取消回调，其目的是clone掉ctx.Done()返回的channel
	ctxParent, _ := context.WithCancel(background)

	ctxChild, cancel := context.WithCancel(ctxParent)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(name string) {
		defer wg.Done()
		select {
		case <-ctxParent.Done():
			fmt.Printf("%s ctx cancel \n", name)
		case <-time.After(2 * time.Second):
			fmt.Printf("%s time after 2 seconds \n", name)
		}
	}("parent")

	wg.Add(1)
	go func(name string) {
		defer wg.Done()
		select {
		case <-ctxChild.Done():
			fmt.Printf("%s ctx cancel \n", name)
		case <-time.After(2 * time.Second):
			fmt.Printf("%s time after 2 seconds \n", name)
		}
	}("child")

	time.Sleep(time.Second)
	cancel()
	wg.Wait()
	fmt.Println("success~~")

}
