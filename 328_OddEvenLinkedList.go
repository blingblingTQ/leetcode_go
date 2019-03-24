package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd, even := head, head.Next
	for even != nil && even.Next != nil {
		tmp := odd.Next
		odd.Next = even.Next
		even.Next = odd.Next.Next
		odd.Next.Next = tmp

		odd = odd.Next
		even = even.Next
	}
	return head
}
