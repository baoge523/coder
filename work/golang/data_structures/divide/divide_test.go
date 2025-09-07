package divide

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/search-a-2d-matrix-ii/
func TestDivide(t *testing.T) {

	arr := []int{1, 2, 3, 5, 7, 11, 19, 20, 34, 77}

	if search(arr, 0, len(arr)-1, 78) {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

}

func search(arr []int, start, end, target int) bool {
	if start > end {
		return false
	}

	middle := start + (end-start)/2
	if arr[middle] == target {
		return true
	} else if arr[middle] > target {
		return search(arr, start, middle-1, target)
	} else {
		return search(arr, middle+1, end, target)
	}
}

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 5, 7, 11, 19, 20, 34, 77}

	target := 3

	low := 0
	high := len(arr) - 1
	isFind := false
	// 这里low == high 表示两个指针指向了同一个索引位置
	for low <= high {
		middle := low + (high-low)/2
		if arr[middle] == target {
			isFind = true
			break
		} else if arr[middle] > target {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	if !isFind {
		fmt.Println("not find")
	} else {
		fmt.Println("find")
	}

}

func TestSearchAsTree(t *testing.T) {
	arr := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 3
	if searchAsTree(arr, target) {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}
}

func searchAsTree(matrix [][]int, target int) bool {
	row := len(matrix)
	high := len(matrix[0])
	i, j := 0, high-1

	for i < row && j >= 0 {
		curr := matrix[i][j]
		if curr == target {
			return true
		} else if curr > target {
			j--
		} else {
			i++
		}
	}
	return false
}
