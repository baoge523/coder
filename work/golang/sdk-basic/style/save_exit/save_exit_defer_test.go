package save_exit

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/*
*

	在该测试样例中，构建了服务a b c，其中 b c依赖于a，同时b c也作为生产者，向chan中投递数据，c作为消费者，消费chan中的数据
	当5秒后，方法结束，调用defer中的方法时，执行顺序为 ccleanup  bcleanup  acleanup
	其中 ccleanup  bcleanup 都是通知后台任务停止向chan中写入数据
	    acleanup 是关闭chan，然后消费完chan里面的数据就退出
*/
func TestExit(t *testing.T) {

	aaa()
	time.Sleep(2 * time.Second)
	fmt.Println("main thread success exit")

}

func aaa() {
	// build channel
	msg := make(chan string, 100)
	ctx := context.Background()
	// build a
	serviceA, acleanup, err := BuildA(nil, "A", msg)
	if err != nil {
		fmt.Println("err")
		return
	}
	defer acleanup()
	serviceA.start(ctx)

	b, bcleanup, err := BuildB(serviceA, "B", msg)
	if err != nil {
		fmt.Println("err")
		return
	}

	defer bcleanup()

	b.start(ctx)

	c, ccleanup, err := BuildC(serviceA, "C", msg)
	if err != nil {
		fmt.Println("err")
		return
	}

	defer ccleanup()

	c.start(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("aaa success")
}

type BuildObjFace interface {
	BuildObj(obj interface{}, name string)
}

type BaseService struct {
	Obj  interface{}
	Name string
	Msg  chan string
}

func (s *BaseService) String() string {
	service := s.Obj.(*BaseService)
	return fmt.Sprintf("base %s , current service name %s", service.Name, s.Name)
}

type ServiceA struct {
	base *BaseService
}

type ServiceB struct {
	*BaseService
	Cancel context.CancelFunc
}

type ServiceC struct {
	*BaseService
	Cancel context.CancelFunc
}

func (s *ServiceA) start(ctx context.Context) {
	fmt.Println("start a consume msg from channel")

	go func() {
		for {
			select {
			case m, ok := <-s.base.Msg: // 注意这种写法
				if !ok {
					fmt.Println("service a msg channel close")
					return
				}
				fmt.Printf("service a consume message: %s \n", m)

			}

		}
	}()
}

func (s *ServiceB) start(ctx context.Context) {
	fmt.Println("start b write msg to channel")

	ctx, s.Cancel = context.WithCancel(ctx)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("start b ctx done ~")
				return
			case <-time.After(time.Second):
				s.Msg <- "service b mock message to channel"
			}
		}
	}()
}

func (s *ServiceC) start(ctx context.Context) {
	fmt.Println("start c write msg to channel")
	ctx, s.Cancel = context.WithCancel(ctx)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("start c ctx done ~")
				return
			case <-time.After(time.Second):
				s.Msg <- "service c mock message to channel"
			}
		}
	}()
}

func BuildA(obj interface{}, name string, msg chan string) (*ServiceA, func(), error) {
	a := &ServiceA{
		base: &BaseService{
			Obj:  obj,
			Name: name,
			Msg:  msg,
		},
	}

	return a, func() {
		close(a.base.Msg)
		fmt.Println("service a close msg channel")
	}, nil
}

func BuildB(obj interface{}, name string, msg chan string) (*ServiceB, func(), error) {
	b := &ServiceB{
		BaseService: &BaseService{
			Obj:  obj,
			Name: name,
			Msg:  msg,
		},
	}
	return b, func() {
		b.Cancel()
		fmt.Println("service b run cancel")
	}, nil
}

func BuildC(obj interface{}, name string, msg chan string) (*ServiceC, func(), error) {
	c := &ServiceC{
		BaseService: &BaseService{
			Obj:  obj,
			Name: name,
			Msg:  msg,
		},
	}
	return c, func() {
		c.Cancel()
		fmt.Println("service c run cancel")
	}, nil
}

func TestName(t *testing.T) {
	background := context.Background()

	ctx, cancelFunc := context.WithCancel(background)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context done")
				return
			}
		}
	}()

	time.Sleep(time.Second)
	defer cancelFunc()

	time.Sleep(time.Second)

}
