package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(arr); i++ {
		go func() {
			fmt.Printf("go run %d\n", i)
		}()
		time.Sleep(1 * time.Millisecond)
		defer fmt.Printf("defer %d\n", i)
	}
	time.Sleep(20 * time.Millisecond)
}
