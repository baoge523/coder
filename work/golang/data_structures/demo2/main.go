package main

import "fmt"

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	ints := merge(nums1, 6, nums2, 3)
	fmt.Println(ints)
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	// 思想：类似于插入排序法的实现
	if m <= n {
		return nums1
	}
	index1 := m - n - 1

	for i := 0; i < n; i++ {
		v := nums2[i]
		j := index1
		for ; j >= 0; j-- {
			if v < nums1[j] {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
			} else {
				break
			}
		}
		nums1[j+1] = v
		index1++
	}
	return nums1
}


