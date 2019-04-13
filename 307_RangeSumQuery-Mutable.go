package main

import (
	"fmt"
)

type SegmentTreeNode struct {
	Start int
	End   int
	Sum   int
	Left  *SegmentTreeNode
	Right *SegmentTreeNode
}

func BuildTree(start int, end int, data []int) *SegmentTreeNode {
	if len(data) == 0 {
		return nil
	}
	root := &SegmentTreeNode{
		Start: start,
		End:   end,
		Sum:   0,
		Left:  nil,
		Right: nil,
	}

	if start == end {
		root.Sum = data[start]
		return root
	}

	mid := start + (end-start)/2
	root.Left = BuildTree(start, mid, data)
	root.Right = BuildTree(mid+1, end, data)
	root.Sum = root.Left.Sum + root.Right.Sum
	return root
}

func UpdateTree(root *SegmentTreeNode, index int, val int) {
	if root == nil {
		return
	}

	if root.Start == root.End {
		root.Sum = val
		return
	}

	mid := root.Start + (root.End-root.Start)/2
	if index <= mid {
		UpdateTree(root.Left, index, val)
	} else {
		UpdateTree(root.Right, index, val)
	}
	root.Sum = root.Left.Sum + root.Right.Sum
	return
}

func QueryTree(root *SegmentTreeNode, start, end int) int {
	if root == nil {
		return 0
	}
	if start == root.Start && end == root.End {
		return root.Sum
	}

	mid := root.Start + (root.End-root.Start)/2
	if mid >= end {
		return QueryTree(root.Left, start, end)
	} else if start > mid {
		return QueryTree(root.Right, start, end)
	} else {
		return QueryTree(root.Left, start, mid) +
			QueryTree(root.Right, mid+1, end)
	}
}

func TraverseTree(root *SegmentTreeNode) {
	if root == nil {
		fmt.Printf("nil ")
		return
	}
	fmt.Printf("[%d, %d, sum=%d] ", root.Start, root.End, root.Sum)
	TraverseTree(root.Left)
	TraverseTree(root.Right)
}

type NumArray struct {
	Root *SegmentTreeNode
}

func Constructor(nums []int) NumArray {
	return NumArray{
		Root: BuildTree(0, len(nums)-1, nums),
	}
}

func (this *NumArray) Update(i int, val int) {
	UpdateTree(this.Root, i, val)
}

func (this *NumArray) SumRange(i int, j int) int {
	return QueryTree(this.Root, i, j)
}

func (this *NumArray) Debug() {
	TraverseTree(this.Root)
	fmt.Println()
}

func main() {
	nums := []int{9, -8}
	numArr := Constructor(nums)
	numArr.Debug()

	numArr.Update(0, 3)
	numArr.Debug()
	fmt.Println(numArr.SumRange(1, 1))
	numArr.Debug()
	fmt.Println(numArr.SumRange(0, 1))
	numArr.Debug()
	numArr.Update(1, -3)
	numArr.Debug()
	fmt.Println(numArr.SumRange(0, 1))
	numArr.Debug()
}
