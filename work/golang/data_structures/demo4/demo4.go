package main

import "fmt"

func main() {

	// 9
	arr := []int{1, 11, 5, 9, 2, 77, 22, 66, 14, 8, 9}

	//arr := []int{1, 7, 5, 11, 2, 66, 14, 7, 88}
	fmt.Println(arr)
	heap(arr, 0, len(arr)-1)

	fmt.Println(arr)

}

func quick(arr []int, start, end int) {
	val := arr[end] // 每次最后一个值作为比较值

	tmpStart := start
	tmpEnd := end

	for tmpStart < tmpEnd {

		for tmpStart < tmpEnd {
			curr := arr[tmpStart]
			if curr < val {
				tmpStart++
			} else {
				arr[tmpEnd] = curr
				tmpEnd--
				break // 换个方向执行
			}
		}

		for tmpStart < tmpEnd {
			curr := arr[tmpEnd]
			if curr >= val {
				tmpEnd--
			} else {
				arr[tmpStart] = curr
				tmpStart++
				break
			}
		}
	}

	arr[tmpStart] = val
	if tmpStart > start+1 {
		quick(arr, start, tmpStart-1)
	}
	if tmpStart < end-1 {
		quick(arr, tmpStart+1, end)
	}

}

// 归并排序
func heap(arr []int, start, end int) {
	if start >= end {
		return
	}
	middle := start + (end-start)/2
	heap(arr, start, middle)
	heap(arr, middle+1, end)
	tmp := make([]int, 0, end-start+1)
	tmpStart := start
	tmpMiddle := middle + 1

	for tmpStart <= middle && tmpMiddle <= end {

		if arr[tmpStart] < arr[tmpMiddle] {
			tmp = append(tmp, arr[tmpStart])
			tmpStart++
		} else {
			tmp = append(tmp, arr[tmpMiddle])
			tmpMiddle++
		}
	}
	// 这里是for循环，将所有的值都写入到临时队列中
	for tmpStart <= middle {
		tmp = append(tmp, arr[tmpStart])
		tmpStart++
	}

	// 这里是for循环，将所有的值都写入到临时队列中
	for tmpMiddle <= end {
		tmp = append(tmp, arr[tmpMiddle])
		tmpMiddle++
	}

	for i := 0; i < len(tmp); i++ {
		arr[start+i] = tmp[i]
	}
}
