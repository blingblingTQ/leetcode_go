package main

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	var list MyLinkedList
	return list
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index >= this.Length || index < 0 {
		return -1
	}
	curr := this.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	return curr.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	newNode := Node{
		Val:  val,
		Next: nil,
	}
	newNode.Next = this.Head
	this.Head = &newNode
	if this.Tail == nil {
		this.Tail = &newNode
	}
	this.Length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	newNode := Node{
		Val:  val,
		Next: nil,
	}
	if this.Tail == nil {
		this.Head = &newNode
	} else {
		this.Tail.Next = &newNode
	}
	this.Tail = &newNode
	this.Length++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Length || index < 0 {
		return
	}
	newNode := Node{
		Val:  val,
		Next: nil,
	}

	if index == 0 {
		newNode.Next = this.Head
		this.Head = &newNode
		if this.Tail == nil {
			this.Tail = &newNode
		}
	} else {
		curr := this.Head
		for i := 0; i < index-1; i++ {
			curr = curr.Next
		}
		next := curr.Next
		newNode.Next = next
		curr.Next = &newNode
		if this.Length == index {
			this.Tail = &newNode
		}
	}
	this.Length++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Length {
		return
	}

	if index == 0 {
		this.Head = this.Head.Next
	} else {
		curr := this.Head
		for i := 0; i < index-1; i++ {
			curr = curr.Next
		}
		curr.Next = curr.Next.Next
		if index == this.Length-1 {
			this.Tail = curr
		}
	}
	this.Length--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
