package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func inorderTraverse(root *TreeNode, array *[]int) {
	if root == nil {
		return
	}

	inorderTraverse(root.Left, array)
	*array = append(*array, root.Val)
	inorderTraverse(root.Right, array)
}

func minDiffInBST(root *TreeNode) int {
	array := []int{}
	inorderTraverse(root, &array)
	minDiff := math.MaxInt32
	for i := 0; i < len(array)-1; i++ {
		minDiff = min(minDiff, array[i+1]-array[i])
	}
	return minDiff
}
