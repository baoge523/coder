package main

import "fmt"
import "sync"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	c1 := make(chan struct{})
	c2 := make(chan struct{})

	w := sync.WaitGroup{}

	w.Add(2)

	go func() {
		defer w.Done()

		for i := 0; i < len(nums); {
			<-c1
			fmt.Printf("%d  ", nums[i])
			c2 <- struct{}{}
			i += 1
		}

	}()

	go func() {
		defer w.Done()

		for i := 0; i < len(nums); {
			<-c2
			fmt.Printf("%d  ", nums[i])
			if i+1 < len(nums) {
				c1 <- struct{}{}
			}
			i += 1
		}

	}()

	// start
	c1 <- struct{}{}
	w.Wait()

	close(c1)
	close(c2)

}
