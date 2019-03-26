package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func levelTraverse(root *TreeNode, res *[]int, level int) {
	if root == nil {
		return
	}

	if len(*res) == level {
		*res = append(*res, root.Val)
	} else {
		(*res)[level] = max((*res)[level], root.Val)
	}
	levelTraverse(root.Left, res, level+1)
	levelTraverse(root.Right, res, level+1)
}

func largestValues(root *TreeNode) []int {
	res := []int{}
	levelTraverse(root, &res, 0)
	return res
}
