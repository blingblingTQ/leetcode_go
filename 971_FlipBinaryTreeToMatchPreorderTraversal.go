package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, voyage []int, index *int, res *[]int) {
	if root == nil {
		return
	}

	if root.Val != voyage[*index] {
		*res = []int{-1}
		return
	}

	*index++
	if *index < len(voyage) && root.Left != nil && root.Left.Val != voyage[*index] {
		(*res) = append(*res, root.Val)
		dfs(root.Right, voyage, index, res)
		dfs(root.Left, voyage, index, res)
	} else {
		dfs(root.Left, voyage, index, res)
		dfs(root.Right, voyage, index, res)
	}
}

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	res := []int{}
	index := 0
	dfs(root, voyage, &index, &res)

	if len(res) > 0 && res[0] == -1 {
		return []int{-1}
	}
	return res
}
