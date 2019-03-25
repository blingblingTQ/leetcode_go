package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, curr *ListNode = nil, head
	for curr != nil {
		if curr.Next != nil && curr.Next.Val == curr.Val {
			for curr.Next != nil && curr.Next.Val == curr.Val {
				curr = curr.Next
			}
			if prev == nil {
				head = curr.Next
			} else {
				prev.Next = curr.Next
			}
			curr = curr.Next
		} else {
			prev = curr
			curr = curr.Next
		}
	}
	return head
}
