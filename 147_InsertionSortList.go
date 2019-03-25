package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	sortedEnd := head
	curr := head.Next
	for curr != nil {
		if curr.Val >= sortedEnd.Val {
			sortedEnd = curr
			curr = curr.Next
		} else {
			sortedEnd.Next = curr.Next
			var insertAfter, insertBefore *ListNode = nil, head
			for insertBefore != sortedEnd.Next && insertBefore.Val <= curr.Val {
				insertAfter = insertBefore
				insertBefore = insertBefore.Next
			}
			curr.Next = insertBefore
			if insertAfter == nil {
				head = curr
			} else {
				insertAfter.Next = curr
			}

			curr = sortedEnd.Next
		}
	}

	return head
}
