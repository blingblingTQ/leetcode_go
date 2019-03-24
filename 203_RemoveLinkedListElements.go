package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var dummy ListNode
	dummy.Next = head

	curr := head
	prev := &dummy
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = prev.Next
		}
		curr = curr.Next
	}
	return dummy.Next
}
