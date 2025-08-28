package slice

import (
	"fmt"
	"testing"
)

/*
*

元素数量: 1, 旧容量: 0, 新容量: 1
元素数量: 2, 旧容量: 1, 新容量: 2   // 小于256，翻倍
元素数量: 3, 旧容量: 2, 新容量: 4   // 小于256，翻倍
元素数量: 5, 旧容量: 4, 新容量: 8   // 小于256，翻倍
...
元素数量: 1025, 旧容量: 1024, 新容量: 1280 // 大于256，增加 25% (1024 + 256 * 3/4 = 1024+192=1216，再内存对齐后约为1280)
*/
func main() {
	s := make([]int, 0)
	oldCap := cap(s) // cap 查看slice的容量

	for i := 0; i < 2000; i++ {
		s = append(s, i)
		if newCap := cap(s); newCap != oldCap {
			fmt.Printf("元素数量: %d, 旧容量: %d, 新容量: %d\n", len(s), oldCap, newCap)
			oldCap = newCap
		}
	}
}

func TestDel(t *testing.T) {

	s := make([]int, 0, 10)
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)

	a := s[len(s)-1]
	fmt.Println(a)
	s = s[:len(s)-1] // 左闭右开
	fmt.Println(s)

}
