package main

import "fmt"
import "math"

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

func getMinimumDifference(root *TreeNode) int {
	array := []int{}
	inorderTraverse(root, &array)
	minDiff := math.MaxInt32
	for i := 0; i < len(array)-1; i++ {
		minDiff = min(minDiff, array[i+1]-array[i])
	}
	return minDiff
}

func main() {
	node3 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	node5 := &TreeNode{
		Val:   5,
		Left:  node3,
		Right: nil,
	}
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: node5,
	}
	res := getMinimumDifference(root)
	fmt.Println(res)
}
