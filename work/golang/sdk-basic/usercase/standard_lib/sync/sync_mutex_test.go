package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {

	var m sync.Mutex  // 互斥锁，只能有一个routine got and run

	for i := 0; i < 10; i++ {
		num := i
		go func() {
			fmt.Printf("current %d start \n",num)
			m.Lock()
			defer m.Unlock()
			fmt.Printf("current %d got lock and runing \n",num)
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("main success")

}

// 测试锁的可重入性 -- 不可重入，需要等待释放锁
func TestMutex2(t *testing.T) {
	var m sync.Mutex  //

	go func() {
		m.Lock()
		defer m.Unlock()
		fmt.Println("current goroutine got lock once")
		gotLock(m)
		fmt.Println("exec gotLock success")
	}()

	time.Sleep(5*time.Second)
	fmt.Println("main success")

}

func gotLock(l sync.Mutex) {
	l.Lock()
	defer l.Unlock()
	fmt.Println("gotLock doing")
}

// double invoke sync.Mutex.lock  -- block
func TestDoubleMutexLock(t *testing.T) {
	var m sync.Mutex
	go func() {
		m.Lock()
		defer m.Unlock()
		fmt.Println("current goroutine got lock once")
		m.Lock() // blocked
		defer m.Unlock()
		fmt.Println("exec gotLock success")
	}()

	time.Sleep(time.Second)
	fmt.Println("main success")
}

func TestTryLock(t *testing.T) {
	var m sync.Mutex
	go func() {
		m.Lock()
		defer m.Unlock()
		fmt.Println("current goroutine got lock once")
		isGot := m.TryLock() // not block
		fmt.Printf("tryLock got %t \n",isGot)
		if isGot {
			defer m.Unlock()
			fmt.Println("current goroutine got lock again")
		}

		fmt.Println("goroutine run success")
	}()

	time.Sleep(time.Second)
	fmt.Println("main success")
}

// 使用位运算是最快的计算方式
func TestCompare(t *testing.T) {
	old := 1
	mutexLocked :=1
	mutexStarving := 4

	fmt.Println(mutexLocked|mutexStarving) // | 表示二进制的或操作，有1就为1： 001 | 100 = 101 ==> 1|4=5

	// 0&(1|4) == 1  ===> no
	// 1&(1|4) == 1  ===> ok
	if old&(mutexLocked|mutexStarving) == mutexLocked {
		fmt.Printf("when old = %d, old&(mutexLocked|mutexStarving) == mutexLocked \n",old)
	}
	fmt.Println("success")

}
