package main

import (
	"fmt"
	"testing"
)

type NodeN struct {
	Num  int
	Next *NodeN
	// Pre  *NodeN
}

func initNoden() *NodeN {
	num5 := &NodeN{Num: 5, Next: nil}
	num4 := &NodeN{Num: 4, Next: num5}
	num3 := &NodeN{Num: 3, Next: num4}
	num2 := &NodeN{Num: 2, Next: num3}
	root := &NodeN{Num: 1, Next: num2}
	return root
}
func reversal(root *NodeN) *NodeN {
	var next *NodeN

	for root != nil {
		temp := root
		root = root.Next
		temp.Next = next
		next = temp
	}
	return next
}
func print(root *NodeN) {
	tmp := root
	for tmp != nil {
		fmt.Println(tmp.Num)
		tmp = tmp.Next
	}
}

func TestNode(t *testing.T) {
	root := initNoden()
	print(root)
	root = reversal(root)
	fmt.Println("---------")
	print(root)
}

func Test1(t *testing.T) {
	s := " "
	m := make(map[byte]int) // 这里可以根据字符串的特性，比如根据asicc的特性验证

	maxLen := 0
	left := 0
	right := 0
	for ; right < len(s); right++ {
		curr := s[right]
		if _, ok := m[curr]; ok {
			if right-left > maxLen {
				maxLen = right - left
			}
			if left < m[curr]+1 {
				left = m[curr] + 1
			}
		}
		m[curr] = right
	}
	// 走完后，需要再判断一次
	if right-left > maxLen {
		maxLen = right - left
	}
	fmt.Println(maxLen)
}
