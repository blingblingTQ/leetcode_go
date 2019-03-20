package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTreeCore(preorder []int, preStart int, preEnd int,
	inorder []int, inStart int, inEnd int) *TreeNode {
	if preStart > preEnd || preEnd-preStart != inEnd-inStart {
		return nil
	}

	rootVal := preorder[preStart]
	root := &TreeNode{
		Val:   rootVal,
		Left:  nil,
		Right: nil,
	}

	var index int
	for index = inStart; index <= inEnd; index++ {
		if inorder[index] == rootVal {
			break
		}
	}
	if index > inEnd {
		//error, do not find
		return nil
	}
	root.Left = buildTreeCore(preorder, preStart+1, preStart+index-inStart, inorder, inStart, index-1)
	root.Right = buildTreeCore(preorder, preStart+index-inStart+1, preEnd, inorder, index+1, inEnd)
	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildTreeCore(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}
