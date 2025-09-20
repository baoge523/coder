package array_exe

import (
	"fmt"
	"testing"
)

/*
*
找一个数组中的连续值等于指定值的的情况
14 目标值
{1,2,6,8,2,5,7} 给定数组
输出：
{{6,8},{2,5,7}}
*/
func Test1(t *testing.T) {
	arr := []int{1, 2, 6, 8, 2, 5, 7, 2, 9, 5, 11, 5, 6, 3}
	val := 14

	var target [][]int
	left := 0
	right := 0
	total := 0
	for right < len(arr) {
		if total == val {
			var tmp []int
			// 这里需要注意是先比较，在加到total中，所有比较的total是不包含右指针的，所以这里需要小于右指针
			for tempJ := left; tempJ < right; tempJ++ {
				tmp = append(tmp, arr[tempJ])
			}
			target = append(target, tmp)
			total -= arr[left]
			left++
			continue
		} else if total > val { // 区间值大于val时，左指针移动，且从总和中减掉数据
			total -= arr[left]
			left++
			continue
		}
		total += arr[right] // 累加到总和中
		right++
	}

	// 保证最后一次可以正常加入到数组中
	if total == val {
		var tmp []int
		for tempJ := left; tempJ < right; tempJ++ {
			tmp = append(tmp, arr[tempJ])
		}
		target = append(target, tmp)
	}
	fmt.Println(target)
}
