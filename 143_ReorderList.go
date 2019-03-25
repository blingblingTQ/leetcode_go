package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func mergeList(before *ListNode, after *ListNode) *ListNode {
	if before == nil {
		return after
	}
	if after == nil {
		return before
	}

	behind := mergeList(before.Next, after.Next)
	before.Next = after
	after.Next = behind
	return before
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head.Next
	//分成两个链表
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	beforeHead := head
	afterHead := slow.Next
	slow.Next = nil

	//反转后半链表
	revAfterHead := reverseList(afterHead)

	//合并
	mergeList(beforeHead, revAfterHead)
}
