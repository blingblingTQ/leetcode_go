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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	} else if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	} else {
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	}
}


/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    int minDepth(TreeNode* root) {
        if(root == NULL) {
            return 0;
        }

        queue<TreeNode *> q;
        q.push(root);
        int depth = 0;
        int minDepth = INT_MAX;
        while(!q.empty()) {
            depth++;
            int levelSize = q.size();
            for(int i = 0; i < levelSize; i++) {
                TreeNode *node = q.front();
                q.pop();
                if(node->left == NULL && node->right == NULL) {
                    minDepth = min(minDepth, depth);
                } else {
                    if (node->left != NULL) {
                        q.push(node->left);
                    }
                    if (node->right != NULL) {
                        q.push(node->right);
                    }
                }
            }
        }
        return minDepth;
    }
};
