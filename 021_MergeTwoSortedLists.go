package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//iterative
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}

	curr := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
			curr = curr.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
			curr = curr.Next
		}
	}
	if l1 == nil {
		curr.Next = l2
	} else if l2 == nil {
		curr.Next = l1
	}
	return dummy.Next
}

//recursive
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
