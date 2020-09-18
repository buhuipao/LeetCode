/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if isLeave(root.Left) {
        return root.Left.Val + sumOfLeftLeaves(root.Right)
    }
    return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

func isLeave(root *TreeNode) bool {
    return root != nil && root.Left == nil && root.Right == nil
}
