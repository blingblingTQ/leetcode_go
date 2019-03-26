package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelTraverse(root *TreeNode, maxLevel *int, leftMostValue *int, level int) {
	if root == nil {
		return
	}
	if level == *maxLevel {
		*maxLevel++
		*leftMostValue = root.Val
	}
	levelTraverse(root.Left, maxLevel, leftMostValue, level+1)
	levelTraverse(root.Right, maxLevel, leftMostValue, level+1)
}

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0 //error
	}

	maxLevel := 0
	leftMostValue := 0
	levelTraverse(root, &maxLevel, &leftMostValue, 0)
	return leftMostValue
}
