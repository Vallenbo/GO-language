package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}

	return dummy.Next
}

func main() {
	point := &ListNode{Next: nil}
	head1 := point
	fmt.Println("请输入l1:")
	for {
		n, _ := fmt.Scanln(&point.Val)
		fmt.Println("输入的个数为:", n)
		if n == 0 {
			point.Next = nil
			break
		}
		node1 := ListNode{Next: nil}
		point.Next = &node1
		point = &node1
	}

	point2 := &ListNode{Next: nil}
	head2 := point2
	fmt.Println("请输入l2:")
	for {
		n, _ := fmt.Scanln(&point2.Val)
		fmt.Println("输入的个数为:", n)
		if n == 0 {
			point2.Next = nil
			break
		}
		node1 := ListNode{Next: nil}
		point2.Next = &node1
		point2 = &node1
	}

	point = mergeTwoLists(head1, head2)
	fmt.Println("输出为：")
	for point.Next != nil {
		fmt.Print(point.Val, "->")
		point = point.Next
	}
}