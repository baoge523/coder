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
		11, 12,
	}
	copy(copys, s)
	fmt.Println(copys)

	fmt.Println("------")
	fmt.Printf("%p\n", s)
	fmt.Printf("%p\n", copys)

}

func TestSliceCopy2(t *testing.T) {
	arr1 := make([]int, 0, 8)
	arr1 = append(arr1, 1)
	arr1 = append(arr1, 2)
	arr1 = append(arr1, 3)
	arr1 = append(arr1, 4)

	arr2 := arr1 // 共用同一个数组内存，会相互影响

	arr3 := make([]int, len(arr1)) // 创建新内存空间，不会相互影响
	copy(arr3, arr1)
	arr1 = append(arr1, 5)
	arr1 = append(arr1, 6)
	arr1[0] = 222
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}
