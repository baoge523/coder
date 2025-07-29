package context

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {

	// backgroundCtx 继承 emptyCtx
	background := context.Background()

	// cancelCtx
	// cancelFunc 执行时，关闭 ctx.Done() 的chan
	ctx, cancelFunc := context.WithCancel(background)
	defer cancelFunc()

	// valueCtx 其中的属性: context 存放parent  key 存放key值  value 存放value值
	ctx = context.WithValue(ctx, "hello", "world")

	// timerCtx 继承 cancelCtx
	//     1、如果parent.DeadLine() < d(当前设置的时间)；说明parent的早，说明就可以不用创建一个timerCtx对象，返回一个cancelCtx即可
	//     2、通过继承的 cancelCtx 设置ctx的父子关系，父取消时，需要取消子
	//     3、如果设置的时间比当前时间小，那么执行cancel
	//     4、通过 time.AfterFunc 设置定时器，定时执行cancel，同时返回cancel方法，可以手动取消
	// cancelFunc 手动取消，即关闭ctx.Done()
	// 底层调用: WithDeadlineCause(parent, d, nil)
	ctx, cancelFunc = context.WithDeadline(ctx, time.Now().Add(5*time.Second))

	// 底层调用 WithDeadline(parent, time.Now().Add(timeout)
	ctx, cancelFunc = context.WithTimeout(ctx, time.Second)

	// 当定时取消时，可以通过context.Cause(ctx) 获取到cause信息
	ctx, cancelFunc = context.WithDeadlineCause(ctx, time.Now().Add(6*time.Second), context.DeadlineExceeded)
	// 获取到ctx中的cause信息
	err := context.Cause(ctx)
	if errors.Is(err,context.DeadlineExceeded) {

	}

	// 只返回ctx，不返回cancel方法, 只有.value能用，其他的都是空方法
	withoutCancelCtx := context.WithoutCancel(ctx)
	withoutCancelCtx.Value("aaaa")

	// 底层调用 WithDeadlineCause(parent, time.Now().Add(timeout), cause)
	ctx, cancelFunc = context.WithTimeoutCause(ctx, time.Second, context.DeadlineExceeded)

	var customCancel  context.CancelFunc
	customCancel = func() {
		fmt.Println("customCancel ~~~")
	}
	defer customCancel()

	var customCancelCause context.CancelCauseFunc
	customCancelCause = func(cause error) {
		fmt.Println(cause)
	}
	defer customCancelCause(context.Canceled)


}

