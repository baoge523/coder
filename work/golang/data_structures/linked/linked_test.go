package linked

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	sb := strings.Builder{}
	temp := ln
	for temp != nil {
		sb.WriteString(strconv.Itoa(temp.Val))
		sb.WriteString(",")
		temp = temp.Next
	}
	return sb.String()
}

func TestName(t *testing.T) {
	node4 := &ListNode{
		Val: 4,
	}
	node3 := &ListNode{
		Val:  3,
		Next: node4,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node3,
	}
	root := &ListNode{
		Val:  1,
		Next: node2,
	}

	node22 := &ListNode{
		Val:  22,
		Next: node3,
	}
	root11 := &ListNode{
		Val:  11,
		Next: node22,
	}
	fmt.Println(root)
	fmt.Println(root11)
	node := getIntersectionNode(root, root11)
	if node != nil {
		fmt.Println(node.Val)
	}

	fmt.Println(root)
	fmt.Println(root11)

}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var target *ListNode
	// 翻转
	reverseHeadeA := reverseList(headA)
	fmt.Println("--------")
	fmt.Println(reverseHeadeA)

	reverseHeadeB := reverseList(headB)
	fmt.Println(reverseHeadeB)

	//// 说明有相交点
	//if reverseHeadeA == reverseHeadeB {
	//	tempA := reverseHeadeA
	//	tempB := reverseHeadeB
	//	for tempA == tempB && tempA != nil && tempB != nil {
	//		target = tempA
	//		tempA = tempA.Next
	//		tempB = tempB.Next
	//	}
	//}

	// 翻转回去
	headA = reverseList(reverseHeadeA)
	headB = reverseList(reverseHeadeB)
	return target
}

func reverseList(head *ListNode) *ListNode {

	var target *ListNode
	for head != nil {
		temp := head.Next
		head.Next = target
		target = head
		head = temp
	}

	return target
}
