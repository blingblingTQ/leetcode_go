package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getLen(head *ListNode) int {
	length := 0
	for curr := head; curr != nil; curr = curr.Next {
		length++
	}
	return length
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 || k == 1 {
		return head
	}
	length := getLen(head)
	time := length / k

	nextStart := head
	start := head
	var prevLast, retHead *ListNode = nil, head
	for i := 0; i < time; i++ {
		start = nextStart
		curr := nextStart
		for j := 0; j < k-1; j++ {
			curr = curr.Next
		}
		nextStart = curr.Next
		curr.Next = nil
		reverseList(start)
		if prevLast == nil {
			retHead = curr
		} else {
			prevLast.Next = curr
		}
		start.Next = nextStart
		prevLast = start
	}
	return retHead
}
