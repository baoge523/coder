package sync_pool

import (
	"fmt"
	"sync"
	"test/share/entity"
	"testing"
)

// sync.pool的目的，就是复用对象，该复用的主要是生命周期比较长的且是无状态的，如果是short-live的对象，不建议使用sync.Pool
// 参考fmt中的使用了sync.Pool来存放pp，达到复用

// sync.Pool.Get 有选择性的忽略pool而是用New;所以存放的对象必须是无状态的
var UserPool = sync.Pool{
	New: func() any {
		return &entity.User{}
	},
}

func TestSyncPool(t *testing.T) {
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	user := UserPool.Get().(*entity.User)
	if user.Name == "" {
		user.Name = fmt.Sprintf("zhangsan")
	}
	UserPool.Put(user)
	fmt.Printf("begin:  user name = %s \n", user.Name)
	for i := 0; i < 5; i++ {
		num := i
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			user := UserPool.Get().(*entity.User) // 把数据拿走后，当前pool就没有数据了，此时来拿，就会返回new数据
			fmt.Printf("before: current goroutine num = %d user name = %s \n", num, user.Name)
			if user.Name == "" {
				user.Name = fmt.Sprintf("zhangsan_%d", num)
			}
			UserPool.Put(user)
			fmt.Printf("after: current goroutine num = %d user name = %s \n", num, user.Name)
		}(num)
	}
	wg.Wait()
	fmt.Println("success")
	user = UserPool.Get().(*entity.User)
	fmt.Println(user)
	UserPool.Put(user)
}

type A struct {
	buf []byte
	m   map[string]string
}

func Test1(t *testing.T) {

	//a := new(A)
	//buf := &a.buf
	//fmt.Println(buf)
	//b := A{}
	//buf1 := b.buf
	//fmt.Println(buf1)

	a := new(A)
	m := &a.m // 获取的变量的地址，不是变量指定的内存地址
	a.m = make(map[string]string)

	(*m)["aa"] = "aa"
	fmt.Println(*m)

}
