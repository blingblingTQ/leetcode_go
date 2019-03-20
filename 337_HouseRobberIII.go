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

func robCore(root *TreeNode, cache map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if _, ok := cache[root]; ok {
		return cache[root]
	}

	noChooseMax := robCore(root.Left, cache) + robCore(root.Right, cache)

	chooseMax := root.Val
	if root.Left != nil {
		chooseMax += robCore(root.Left.Left, cache) + robCore(root.Left.Right, cache)
	}
	if root.Right != nil {
		chooseMax += robCore(root.Right.Left, cache) + robCore(root.Right.Right, cache)
	}
	cache[root] = max(noChooseMax, chooseMax)
	return cache[root]
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cache := make(map[*TreeNode]int)
	return robCore(root, cache)
}
