package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev, curr := head, head.Next
	for curr != nil {
		if curr.Val != prev.Val {
			prev.Next = curr
			prev = curr
		}
		curr = curr.Next
	}
	prev.Next = nil
	return head
}
