package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//反转链表法
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	fast, slow := head, head
	var slowPrev, slowNext *ListNode = nil, nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slowNext = slow.Next
		slow.Next = slowPrev
		slowPrev = slow
		slow = slowNext
	}

	var left, right *ListNode = nil, nil
	if fast == nil {
		left = slowPrev
		right = slow
	} else {
		left, right = slowPrev, slow.Next
	}
	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

// 递归法
func isPalindromeCore(head *ListNode, length int) (bool, *ListNode) {
	if head == nil || length == 0 {
		return true, nil
	} else if length == 1 {
		return true, head.Next
	} else if length == 2 {
		return head.Val == head.Next.Val, head.Next.Next
	} else {
		res, next := isPalindromeCore(head.Next, length-2)
		if res == false || next == nil {
			return false, nil
		}
		res = (head.Val == next.Val)
		return res, next.Next
	}
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	length := 0
	for curr := head; curr != nil; curr = curr.Next {
		length++
	}

	result, _ := isPalindromeCore(head, length)
	return result
}

func main() {
	var head *ListNode = nil
	fmt.Println(isPalindrome(head))
}
