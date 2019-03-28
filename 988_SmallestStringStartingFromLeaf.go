package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, ans *string, currPath []byte) {
	if root == nil {
		return
	}

	newPath := append([]byte{byte(root.Val + 'a')}, currPath...)
	if root.Left == nil && root.Right == nil {
		if string(newPath) < *ans || *ans == "" {
			*ans = string(newPath)
		}
		return
	}
	dfs(root.Left, ans, newPath)
	dfs(root.Right, ans, newPath)
}

func smallestFromLeaf(root *TreeNode) string {
	ans := ""
	currPath := []byte{}
	dfs(root, &ans, currPath)
	return ans
}
