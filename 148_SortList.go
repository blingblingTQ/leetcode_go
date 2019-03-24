package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getMiddleOfList(head *ListNode) *ListNode {
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := getMiddleOfList(head)
	next := mid.Next
	mid.Next = nil
	return mergeTwoLists(sortList(head), sortList(next))
}

func printList(head *ListNode) {
	for curr := head; curr != nil; curr = curr.Next {
		fmt.Printf("%d->", curr.Val)
	}
	fmt.Printf("\n")
}

func main() {
	var node1, node2, node3, node4 ListNode
	node1.Val = 4
	node1.Next = &node2
	node2.Val = 2
	node2.Next = &node3
	node3.Val = 1
	node3.Next = &node4
	node4.Val = 3
	node4.Next = nil

	head := &node1
	printList(sortList(head))
}
