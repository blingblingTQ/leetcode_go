package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	length := 0
	var curr, end *ListNode = head, nil
	for curr != nil {
		length++
		end = curr
		curr = curr.Next
	}
	k = k % length
	if k == 0 {
		return head
	}

	curr = head
	for i := 0; i < length-k-1; i++ {
		curr = curr.Next
	}
	newHead := curr.Next
	curr.Next = nil
	end.Next = head
	return newHead
}
