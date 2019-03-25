package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  nil,
		Right: nil,
	}

	leftNums := nums[:mid]
	rightNums := nums[mid+1:]

	root.Left = sortedArrayToBST(leftNums)
	root.Right = sortedArrayToBST(rightNums)
	return root
}
