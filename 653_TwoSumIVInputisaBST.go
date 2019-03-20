package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func find(root *TreeNode, k int, m map[int]bool) bool {
	if root == nil {
		return false
	}
	if _, ok := m[k-root.Val]; ok {
		return true
	}
	m[root.Val] = true
	return find(root.Left, k, m) || find(root.Right, k, m)
}

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]bool)
	return find(root, k, m)
}
