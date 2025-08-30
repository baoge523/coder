package main

import (
	"fmt"
	"sync"
)

type Node struct {
	Num   int
	Left  *Node
	Right *Node
}

func initTree() *Node {
	// 1 - 2 3 -4 5 6 7
	num4 := &Node{Num: 4, Left: nil, Right: nil}
	num5 := &Node{Num: 5, Left: nil, Right: nil}
	num6 := &Node{Num: 6, Left: nil, Right: nil}
	num7 := &Node{Num: 7, Left: nil, Right: nil}

	num2 := &Node{Num: 2, Left: num4, Right: num5}
	num3 := &Node{Num: 3, Left: num6, Right: num7}

	root := &Node{Num: 1, Left: num2, Right: num3}
	return root
}

// pre 前序访问 根 左 右
// 1245367
func pre(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d\n", root.Num)
	pre(root.Left)
	pre(root.Right)

}

// mid 中序遍历 左根右
// 4 2 5 1 6 3 7
func mid(root *Node) {
	if root == nil {
		return
	}
	mid(root.Left)
	fmt.Println(root.Num)
	mid(root.Right)
}

// post 后序遍历 左右根
// 4526731
func post(root *Node) {
	if root == nil {
		return
	}
	post(root.Left)
	post(root.Right)
	fmt.Println(root.Num)
}

func main() {
	root := initTree()
	post(root)
}

func aaa() {
	// 交替打印字母和数字
	a := "abcdefg"
	b := "1234567"

	c1 := make(chan struct{})
	c2 := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < len(a); i++ {
			<-c1
			fmt.Printf("%c\n", a[i])
			c2 <- struct{}{}

		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < len(b); i++ {
			<-c2
			fmt.Printf("%c\n", b[i])
			c1 <- struct{}{}

		}

	}()

	// 启动时，默认让c1先执行
	c1 <- struct{}{}

	wg.Wait()
	<-c1
	close(c1)
	close(c2)

}
