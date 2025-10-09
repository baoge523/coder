package main

/**
实现一个抢红包的功能，可以设置总金额和总个数，为了使大家都能满意，需保证每个红包的金额都在平均值左右，最大波动范围可以是平均值的一半，红包金额值可以到分。
1. 输入: 总金额 100，总红包数 11
2. 设计多个测试用例，确认程序的正确性，即总金额、总红包数使用多种组合
*/

// 思路：输出总个数的金额，同时只和等于总金额; 随机数区间，然后总数递减
func main() {

	total := 100.0
	count := 10

	avg := total / count. // 10

	a := avg / 2		//5

	var target []float

	for i:=0;i<count;i++ {
		// 计算随机值（区间）
		if i == count-1 {
			target = append(target,total)
			break
		}
		curr := calc(total,avg,a)
		target = append(target,curr)
		total -=curr
	}


}
// 需要保证在total下分配，同时尽可能的在avg左右，波动是r
func calc(total float,avg float,r float) float {
	// 随机值的上下范围
	max := avg + r. // 15
	min := avg - r  // 5

	if total <= min { // 左闭又开原则
		max := total
		min := 0.01
	}


	var curr float
	for {
		// 随机值 左闭又开原则
		curr = math.Random(min,max)
		if curr < total {
			break
		}

	}

	return curr
}
