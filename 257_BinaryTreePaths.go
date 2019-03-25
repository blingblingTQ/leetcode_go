package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, path string, res *[]string) {
	if root == nil {
		return
	}
	path = path + "->" + strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		path = path[2:]
		*res = append(*res, path)
	}

	helper(root.Left, path, res)
	helper(root.Right, path, res)
}

func binaryTreePaths(root *TreeNode) []string {
	path := ""
	res := []string{}
	helper(root, path, &res)
	return res
}
