package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum, carry int
	var head ListNode
	var prev = &head
	for l1 != nil && l2 != nil {
		sum = (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10
		newNode := new(ListNode)
		newNode.Val = sum
		prev.Next = newNode
		prev = newNode
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		sum = (l1.Val + carry) % 10
		carry = (l1.Val + carry) / 10
		newNode := new(ListNode)
		newNode.Val = sum
		prev.Next = newNode
		prev = newNode
		l1 = l1.Next
	}
	for l2 != nil {
		sum = (l2.Val + carry) % 10
		carry = (l2.Val + carry) / 10
		newNode := new(ListNode)
		newNode.Val = sum
		prev.Next = newNode
		prev = newNode
		l2 = l2.Next
	}
	if carry != 0 {
		newNode := new(ListNode)
		newNode.Val = carry
		prev.Next = newNode
		prev = newNode
		l2 = l2.Next
	}
	return head.Next
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%d->", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func main() {
	var l13 = &ListNode{
		Val:  3,
		Next: nil,
	}
	var l12 = &ListNode{
		Val:  4,
		Next: l13,
	}
	var l11 = &ListNode{
		Val:  2,
		Next: l12,
	}

	var l23 = &ListNode{
		Val:  4,
		Next: nil,
	}
	var l22 = &ListNode{
		Val:  6,
		Next: l23,
	}
	var l21 = &ListNode{
		Val:  5,
		Next: l22,
	}

	printList(l11)
	printList(l21)
	l := addTwoNumbers(l11, l21)
	printList(l)
}
