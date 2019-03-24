package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//calculate length
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	length1, length2 := 0, 0
	for curr := headA; curr != nil; curr = curr.Next {
		length1++
	}
	for curr := headB; curr != nil; curr = curr.Next {
		length2++
	}
	a, b := headA, headB
	if length1 < length2 {
		for i := 0; i < length2-length1; i++ {
			b = b.Next
		}
	} else {
		for i := 0; i < length1-length2; i++ {
			a = a.Next
		}
	}
	for a != nil && b != nil {
		if a == b {
			return a
		}
		a = a.Next
		b = b.Next
	}
	return nil
}

//better solution
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
		if p1 == nil && p2 != nil {
			p1 = headB
		} else if p1 != nil && p2 == nil {
			p2 = headA
		}
	}
	return p1
}
