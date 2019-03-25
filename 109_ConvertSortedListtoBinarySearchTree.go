package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMiddle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}
	if prev != nil {
		prev.Next = nil
	}
	return slow
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	mid := getMiddle(head)
	if mid == nil {
		return nil
	}

	root := &TreeNode{
		Val:   mid.Val,
		Left:  nil,
		Right: nil,
	}
	if mid == head {
		root.Left = nil
	} else {
		root.Left = sortedListToBST(head)
	}
	root.Right = sortedListToBST(mid.Next)
	return root
}
