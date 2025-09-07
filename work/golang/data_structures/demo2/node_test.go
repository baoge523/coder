package main

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {

	node2 := &ListNode{
		Val: 2,
	}
	head := &ListNode{
		Val:  1,
		Next: node2,
	}
	h, tail := reverseSingle(head)
	fmt.Println(h)
	fmt.Println(tail)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseSingle(head *ListNode) (*ListNode, *ListNode) {
	var target *ListNode
	tail := head
	for head != nil {
		temp := head.Next
		head.Next = target
		target = head
		head = temp
	}

	return target, tail
}

func Test1(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}
	var target ListNode
	hand(head, &target)
	fmt.Println(target)
}

func hand(root *ListNode, target *ListNode) {
	target = root
	fmt.Println(target)
}
