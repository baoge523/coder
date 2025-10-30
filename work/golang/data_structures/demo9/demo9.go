package demo9

/**
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
你返回所有和为 0 且不重复的三元组。

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
*/

// 方式1: 暴力求解： 三重循环判断
// 方式2: 动态规划
func main() {

}

func method1(arr []int) {

	var target [][]int
	arrLen := len(arr)
	i := 0
	j := 1
	k := 2

	for i <= arrLen-3 {
		for j <= arrLen-2 {
			for k <= arrLen-1 {
				if arr[i]+arr[j]+arr[k] == 0 { // 需要去重
					target = append(target, []int{arr[i], arr[j], arr[k]})
				}
				k++
			}
			j++
		}
		i++
	}
}
