package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//interative
func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev, next *ListNode = nil, nil
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

//recursive
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
