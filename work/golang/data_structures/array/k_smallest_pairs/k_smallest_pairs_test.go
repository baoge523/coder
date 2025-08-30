package k_smallest_pairs

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

// 两个有序数组，求n个最小和的组合

// 两种方式： 暴力求解   优化方式

var (
	array1 = []int{1, 7, 11}
	array2 = []int{2, 4, 6}
)

/*
*
说一下这个问题
空间复杂度： O(N)
时间复杂度:  O(N^2)

同时还依赖与排序操作

暴力求解法：
1、将所有可能的结果都存放到数组中
2、按数组排序
3、取前n个
*/
func Test1(t *testing.T) {

	numbers := make(Numbers, 0, len(array1)*len(array2))

	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			sum := array1[i] + array2[j]
			numbers = append(numbers, &Number{Sum: sum, I: array1[i], J: array2[j]})
		}
	}

	sort.Sort(numbers)
	n := 4
	for i := 0; i < n; i++ {
		fmt.Println(numbers[i])
	}
	fmt.Println("hello")
}

// 优化上面的操作  本质就是取前n个最小的数(组合)  -- 可以基于小顶堆实现
// 两个有序数组的组合，构成了一张图，默认从右边压入数据，当j=0的时候，也压入下一行数据
// 通过小顶堆的方式实现比较压入的数据取最小值，最后直到数据取完

// 收获点: golang中实现一个小顶堆 或者 大顶堆 通过container的heap实现的，

// 时间复杂度 O(n logn)
func Test2(t *testing.T) {

	n := 8
	if n <= 0 || n > len(array1)*len(array2) {
		return
	}
	var nums []Number

	i, j := 0, 0

	h := &MinHeap{}

	heap.Push(h, Number{
		Sum:    array1[i] + array2[j],
		I:      array1[i],
		J:      array2[j],
		Indexi: i,
		Indexj: j,
	})

	/**
	思路是: 因为是有序数组，所有最小值是array1[0],array2[0] 后面小的就是临近的，比如array1[0],array2[1]; array1[1],array2[0]
	       但这里就会有重复弹入的风险，所以需要制定一个规则，每行的数据依次弹入，然后遇到J=0，就填入下一行的数据
	   i 表示行
	   j 表示列
	*/

	// 弹出所有的堆中元素
	for h.Len() != 0 {
		min := heap.Pop(h)
		minNum := min.(Number)
		// 保存结果
		nums = append(nums, minNum)
		// 提前退出
		if len(nums) == n {
			break
		}
		if minNum.Indexi < len(array1)-1 {
			heap.Push(h, Number{
				Sum:    array1[minNum.Indexi+1] + array2[minNum.Indexj],
				I:      array1[minNum.Indexi+1],
				J:      array2[minNum.Indexj],
				Indexi: minNum.Indexi + 1,
				Indexj: minNum.Indexj,
			})
		}

		if minNum.Indexi == 0 && minNum.Indexj < len(array2)-1 {
			heap.Push(h, Number{
				Sum:    array1[0] + array2[minNum.Indexj+1],
				I:      array1[0],
				J:      array2[minNum.Indexj+1],
				Indexi: 0,
				Indexj: minNum.Indexj + 1,
			})
		}

	}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

}

func Test3(t *testing.T) {
	num := []int{3, 5, 1, 11, 7, 4}
	h := &MinHeap{}
	heap.Init(h)
	for i := 0; i < len(num); i++ {
		heap.Push(h, Number{
			Sum:    num[i],
			I:      0,
			J:      0,
			Indexi: 0,
			Indexj: 0,
		})
	}

	for h.Len() != 0 {
		min := heap.Pop(h)
		minNum := min.(Number)
		fmt.Println(minNum.Sum)
	}

}

type MinHeap []Number

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i].Sum <= m[j].Sum // 小顶堆
}

// Swap swaps the elements with indexes i and j.
func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(Number))
}

func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

type Number struct {
	Sum    int
	I      int
	J      int
	Indexi int
	Indexj int
}

func (n Number) String() string {
	return fmt.Sprintf("[%d,%d]", n.I, n.J)
}

type Numbers []*Number

func (n Numbers) Len() int {
	return len(n)
}

func (n Numbers) Less(i, j int) bool {
	if n[i].Sum <= n[j].Sum {
		return true
	}
	return false
}

// Swap swaps the elements with indexes i and j.
func (n Numbers) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
