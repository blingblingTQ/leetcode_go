package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, G []int) int {
	m := make(map[int]bool)
	for _, g := range G {
		m[g] = true
	}

	curr := head
	result := 0
	for curr != nil {
		if m[curr.Val] == true {
			for curr != nil && m[curr.Val] == true {
				curr = curr.Next
			}
			result++
		} else {
			curr = curr.Next
		}
	}
	return result
}
