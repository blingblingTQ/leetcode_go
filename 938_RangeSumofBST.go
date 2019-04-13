package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(root *TreeNode, L int, R int, sum *int) {
	if root == nil {
		return
	}

	if root.Val >= L && root.Val <= R {
		*sum += root.Val
	}
	traverse(root.Left, L, R, sum)
	traverse(root.Right, L, R, sum)
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	traverse(root, L, R, &sum)
	return sum
}
