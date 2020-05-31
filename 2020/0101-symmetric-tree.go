/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return root == nil || helper(root.Left, root.Right)
}

func helper(root1, root2 *TreeNode) bool {
    if root1 == nil && root2 == nil {
        return true
    }
    if root1 == nil || root2 == nil {
        return false
    }
    if root1.Val != root2.Val {
        return false
    }
    return helper(root1.Left, root2.Right) && helper(root1.Right, root2.Left)
}
