package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelTraverse(root *TreeNode, res *[]int, level int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, root.Val)
	}
	levelTraverse(root.Right, res, level+1)
	levelTraverse(root.Left, res, level+1)
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	levelTraverse(root, &res, 0)
	return res
}
