package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//fast and slow pointer
func detectCycle1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	hasCycle := false
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			hasCycle = true
			break
		}
	}

	if hasCycle == false {
		return nil
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

// hashmap
func detectCycle2(head *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	curr := head
	for curr != nil {
		if _, ok := m[curr]; ok {
			return curr
		}
		m[curr] = true
		curr = curr.Next
	}
	return nil
}
