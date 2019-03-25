package main
  
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

//如果平衡，返回高度
func checkAndGet(root *TreeNode) (bool, int) {
    if root == nil {
        return true, 0
    }
    balance, leftHeight := checkAndGet(root.Left)
    if !balance {
        return false, -1
    }
    balance, rightHeight := checkAndGet(root.Right)
    if !balance {
        return false, -1
    }
    
    if leftHeight-rightHeight < -1 || leftHeight-rightHeight > 1 {
        retuarn false, -1
    }
    return true, max(leftHeight, rightHeight)+1
}


func isBalanced(root *TreeNode) bool {
    if balance, _ := checkAndGet(root); !balance {
        return false
    }
    return true
}

