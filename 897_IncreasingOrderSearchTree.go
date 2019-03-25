package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := increasingBST(root.Left)
	right := increasingBST(root.Right)

	node := left
	for node != nil && node.Right != nil {
		node = node.Right
	}
	if node != nil {
		node.Right = root
	}
	root.Right = right
	root.Left = nil

	if left == nil {
		return root
	}
	return left
}
