package runtime

import (
	"fmt"
	"sync"
	"testing"
)

// 设置指定goroutine数量去处理job

// 通过指定协程共享job chan 的方式
func TestShareJobs(t *testing.T){
	var group sync.WaitGroup
	jobNum := 10
	threadNum := 3
	// 注意chan需要缓存
	jobs := make(chan string,jobNum)
	results := make(chan string,jobNum)
	for i := 0; i < threadNum; i++ {
		group.Add(1)
		num := i
		go func() {
			defer group.Done()
			for job := range jobs {
				results <- fmt.Sprintf("job:%s,num=%d",job,num)
			}
		}()
	}
	for i := 0; i < jobNum; i++ {
		jobs <- fmt.Sprintf("current job num = %d",i)
	}
	close(jobs)
	group.Wait()
	close(results)
	for result:= range results {
		fmt.Println(result)
	}
}

// 通过信号量的方式限制同时只能有指定个协程数处理
func TestShareSemaphore(t *testing.T){
	jobNum := 10
	threadNum := 3

	var group sync.WaitGroup

	results := make(chan string,jobNum)
	semaphore := make(chan struct{},threadNum)

	for i := 0; i < jobNum; i++ {
		semaphore <- struct{}{}
		group.Add(1)
		num := i
		go func() {
			defer func() {
				<-semaphore
				group.Done()
			}()
			results <- fmt.Sprintf("current job num = %d",num)
		}()
	}

	group.Wait()
	close(semaphore)
	close(results)
	for result := range results {
		fmt.Println(result)
	}


}
