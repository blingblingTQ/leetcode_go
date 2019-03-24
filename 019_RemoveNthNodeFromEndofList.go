package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n == 0 {
		return head
	}
	var prev *ListNode = nil
	fast, slow := head, head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		prev = slow
		slow = slow.Next
	}
	//delete slow
	if prev != nil {
		prev.Next = slow.Next
		return head
	} else {
		return head.Next
	}
}
