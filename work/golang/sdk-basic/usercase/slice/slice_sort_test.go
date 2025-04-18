package slice

import (
	"fmt"
	"sort"
	"testing"
)

func TestSliceSort(t *testing.T) {
	var nums []int
	nums = append(nums, 3)
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 4)
	fmt.Printf("%v \n", nums)
	fmt.Println("---sort---")
	sort.Ints(nums)
	fmt.Printf("%v \n", nums)

}
