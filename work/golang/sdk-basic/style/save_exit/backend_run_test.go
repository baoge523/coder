package save_exit

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestBG(t *testing.T) {
	run()
	time.Sleep(2 * time.Second)
	fmt.Println("main success")
}

func run() {
	task := &BGTask{
		Group:   sync.WaitGroup{},
		ErrChan: make(chan error, 100),
	}
	task.Start(context.Background())
	defer task.Stop()
	go func() {
		for e := range task.ErrChan {
			fmt.Printf("show err: %s \n", e.Error())
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("run success")
}

// 后台任务如何安全退出
type RunningAtBackground interface {
	// start
	Start(ctx context.Context)
	// s(top
	Stop()
	// error 返回一个接收chan error的 chan
	Errors() <-chan error
}

type BGTask struct {
	Cancel  context.CancelFunc
	Group   sync.WaitGroup
	ErrChan chan error
}

func (s *BGTask) Start(ctx context.Context) {
	fmt.Println("start bg task")

	// s.do

	ctx, s.Cancel = context.WithCancel(ctx)

	s.Group.Add(1)
	go func() {
		defer s.Group.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("task 1 ctx done and exit")
				return
			case <-time.After(time.Second):
				if err := s.do(); err != nil {
					s.ErrChan <- fmt.Errorf("business error %w", err)
				}
			}

		}

	}()
}

func (s *BGTask) Stop() {
	// 关闭掉chan
	defer close(s.ErrChan)
	// 执行后台线程停止操作
	s.Cancel()
	// 等待所有的后台线程都执行完了，就表示stop执行完了
	s.Group.Wait()
}

// 这里的 <-chan error 表示返回出去的chan，不能用于写
func (s *BGTask) Errors() <-chan error {
	return s.ErrChan
}

func (s *BGTask) do() error {
	fmt.Println(" business do ~~~~")
	num := rand.Intn(10)
	if num <= 5 {
		return fmt.Errorf("mock errror")
	}

	return nil
}
