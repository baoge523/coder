package slice

import (
	"fmt"
	"testing"
)

// copy(dest, src) desc 必须有存放src的空间,否则只会拷贝部分src
func TestSliceCopy(t *testing.T) {
	var s []int
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)

	fmt.Println(s)
	//copys := make([]int, len(s))
	// var copys []int  // 没有空间，在copy时不会分配空间
	var copys []int = []int{
		11,12,
	}
	copy(copys, s)
	fmt.Println(copys)

	fmt.Println("------")
	fmt.Printf("%p\n", s)
	fmt.Printf("%p\n", copys)



}
