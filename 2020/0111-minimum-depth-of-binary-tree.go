/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    if root.Left == nil {
        return minDepth(root.Right) + 1
    }
    if root.Right ==  nil {
        return minDepth(root.Left) + 1
    }

    L, R := minDepth(root.Left), minDepth(root.Right)
    if L > R {
        return R + 1
    }
    return L + 1
}
