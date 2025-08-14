package work_pool

import (
	"context"
	"errors"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

// WorkerPool 工作池
type WorkerPool struct {
	workerSize         int
	chanCap            int
	pool               *ants.Pool
	taskChan           chan *taskWrapper
	wg                 sync.WaitGroup
	workerRunningCount int64
	sigTerminal        chan struct{} // 用于表示当前线程池已经关闭了

	wrapPool sync.Pool
}

func NewWorkerPool(workerSize int, cap int) *WorkerPool {
	return &WorkerPool{
		workerSize: workerSize,
		chanCap:    cap,
	}
}

func (wp *WorkerPool) Start(ctx context.Context) error {
	var err error
	wp.pool, err = ants.NewPool(wp.workerSize,
		ants.WithExpiryDuration(time.Minute), // 定期清理无效的worker
		ants.WithPreAlloc(true),              // 为worker预分配内存
		ants.WithNonblocking(false),          // 没有worker时，阻塞
		ants.WithMaxBlockingTasks(20),        // 最大的阻塞任务个数，超过了会报错吗？
	)
	// 复用 taskWrapper
	wp.wrapPool.New = func() any {
		return &taskWrapper{}
	}
	wp.taskChan = make(chan *taskWrapper, wp.chanCap)
	wp.sigTerminal = make(chan struct{})

	if err = wp.loopRun(ctx); err != nil {
		return fmt.Errorf("loop run fail %w", err)
	}
	return err
}

func (wp *WorkerPool) Stop(ctx context.Context) error {
	if wp.taskChan == nil {
		return fmt.Errorf("WorkerPool not start, no need stop")
	}
	close(wp.taskChan)
	close(wp.sigTerminal) // 表示当前线程池已经关闭

	wp.wg.Wait()
	wp.pool.Release() // 关闭ants.pool
	return nil
}

func (wp *WorkerPool) loopRun(ctx context.Context) error {
	// 这里的ctx主要用于输出日志，不做退出和超时控制
	for task := range wp.taskChan {
		wp.wg.Add(1)
		if err := wp.pool.Submit(func() {
			defer func() {
				wp.wg.Done()
				wp.wrapPool.Put(task)                       // 复用
				atomic.AddInt64(&wp.workerRunningCount, -1) // 运行数量减一
			}()
			atomic.AddInt64(&wp.workerRunningCount, 1) // 运行数量加一
			// run job
			wp.execTask(task)

		}); err != nil {
			return err
		}
	}
	return nil
}

func (wp *WorkerPool) execTask(taskWrap *taskWrapper) {

	res := taskWrap.job(taskWrap.ctx)
	defer func() {
		if r := recover(); r != nil {
			res.err = fmt.Errorf("%w: %+v, stack: %s", errors.New("recover panic"), r, string(debug.Stack()))
		}
	}()

	// 将结果信息写回taskWrapper中的result chan中
	if taskWrap.res == nil { // 未初始化，说明不关心结果直接返回
		return
	}

	taskWrap.res <- res
}

// RunJob 将任务提交到线程池中
func (wp *WorkerPool) RunJob(ctx context.Context, waitMs int, job JobTask, res chan<- *JobResult) bool {
	taskW, ok := wp.wrapPool.Get().(*taskWrapper)
	if !ok {
		taskW = &taskWrapper{}
	}
	taskW.res = res
	taskW.job = job
	taskW.ctx = ctx
	// pre check
	select {
	case <-ctx.Done():
		return false
	case <-wp.sigTerminal:
		return false
	default:
		// check ok , go next
	}
	// 队列满时的处理方式

	// 1、一直等待，直到能写入队列
	if waitMs <= -1 {
		select {
		case <-ctx.Done():
			return false
		case wp.taskChan <- taskW:
			return true
		}
	}
	// 2、队列满时，直接抛弃任务
	if waitMs == 0 {
		select {
		case wp.taskChan <- taskW:
			return true
		default:
			return false
		}
	}

	// 3、等待指定时间，如果还不能写入队列，抛弃任务
	select {
	case <-ctx.Done():
		return false
	case wp.taskChan <- taskW:
		return true
	case <-time.After(time.Duration(waitMs) * time.Millisecond):
		return false
	}

}

func (wp *WorkerPool) GetRunningState() WorkerPoolState {
	state := WorkerPoolState{
		WorkerSize:        wp.workerSize,
		WorkerRunningSize: int(atomic.LoadInt64(&wp.workerRunningCount)),
		QueueUse:          len(wp.taskChan),
		QueueCap:          wp.chanCap,
	}
	state.antsRunningSize = wp.pool.Running()
	state.WorkerRunningRate = float32(state.WorkerRunningSize) / float32(state.WorkerSize) * 100
	state.QueueUseRate = float32(state.QueueUse) / float32(state.QueueCap) * 100
	return state
}

type JobResult struct {
	uuid string // 用于关联到某个任务上
	res  interface{}
	err  error
}

// JobTask 定义用户任务模型，用户要执行一个什么样的任务，并返回结果
type JobTask func(ctx context.Context) *JobResult

type taskWrapper struct {
	ctx context.Context   // 当前任务的ctx信息
	job JobTask           // 当前任务
	res chan<- *JobResult // 任务执行结果
}

type WorkerPoolState struct {
	WorkerSize        int
	WorkerRunningSize int
	WorkerRunningRate float32

	QueueCap     int
	QueueUse     int
	QueueUseRate float32

	antsRunningSize int
}
