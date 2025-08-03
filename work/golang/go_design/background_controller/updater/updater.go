package updater

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync/atomic"
	"time"
)

type Updater struct {
	Time         time.Duration
	ValueSetter  func(ctx context.Context) (interface{}, error)
	ErrorHandler func(ctx context.Context, err error)

	// 内部信息
	cancelFunc context.CancelFunc // 用于在stop时，通知内部goroutine不要干活了
	eg         errgroup.Group     // 异步处理任务，同时可以在stop时，执行了cancelFunc后，等待goroutine执行完；
	val        atomic.Value       // 内部数据存储
}

func (r *Updater) Start(ctx context.Context) error {

	// 启动时：初始化
	if err := r.update(ctx); err != nil {
		//
		return err
	}
	ctx, r.cancelFunc = context.WithCancel(ctx)
	// 启动定时任务执行
	r.eg.Go(func() error {
		r.loopRun(ctx)
		return nil
	})

	return nil
}

func (r *Updater) Stop() error {
	r.cancelFunc()     // 通知内部goroutine，被取消了
	return r.eg.Wait() // 等待内部goroutine执行完 -- 重要，这个参数goroutine优雅退出的关键呀
}

// Value 获取值
func (r *Updater) Value() interface{} {
	return r.val.Load()
}

func (r *Updater) update(ctx context.Context) (err error) {
	// 捕获业务可能不可能的panic异常
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("如果遇到的需要recover的错误，转换成可以处理的错误")
		}
	}()

	value, err := r.ValueSetter(ctx)
	if err != nil {
		return err
	}
	r.val.Store(value)
	return nil
}

// loopRun 定时查询业务数据，当stop时取消执行
func (r *Updater) loopRun(ctx context.Context) {
	for {
		select {
		case <-time.After(r.Time):
			if err := r.update(ctx); err != nil && r.ErrorHandler != nil {
				r.ErrorHandler(ctx, err) // 用户自己处理这个错误
			}
		case <-ctx.Done():
			return
		}

	}

}
