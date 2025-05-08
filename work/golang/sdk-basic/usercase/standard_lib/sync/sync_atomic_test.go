package sync

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	num := atomic.Int64{}

	load := num.Load()
	fmt.Println(load)
	newAdd := num.Add(123)
	fmt.Println(newAdd)
	old := num.Swap(456)
	fmt.Printf("old = %d and now = %d \n",old,num.Load())
	if swapped := num.CompareAndSwap(456, 789); swapped {
		fmt.Println("compare and swap success ")
	}
	fmt.Println("success")
}
