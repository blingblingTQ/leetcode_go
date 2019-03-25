package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	var curr, prev *ListNode = head, nil
	for i := 0; i < m-1; i++ {
		prev = curr
		curr = curr.Next
	}

	end := curr
	tmpPrev := curr
	curr = curr.Next
	for i := m + 1; i <= n; i++ {
		next := curr.Next
		curr.Next = tmpPrev
		tmpPrev = curr
		curr = next
	}
	end.Next = curr
	if prev != nil {
		prev.Next = tmpPrev
	} else {
		head = tmpPrev
	}

	return head
}
