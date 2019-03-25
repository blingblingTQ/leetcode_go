package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*res = append(*res, root.Val)
	}
	dfs(root.Left, res)
	dfs(root.Right, res)
}

func getLeafSequence(root *TreeNode) []int {
	res := []int{}
	dfs(root, &res)
	return res
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafs1 := getLeafSequence(root1)
	leafs2 := getLeafSequence(root2)

	if len(leafs1) != len(leafs2) {
		return false
	}

	for i := 0; i < len(leafs1); i++ {
		if leafs1[i] != leafs2[i] {
			return false
		}
	}
	return true
}
