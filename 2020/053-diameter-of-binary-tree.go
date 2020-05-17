/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // 递归实现
func diameterOfBinaryTree(root *TreeNode) int {
    var MAX int
    helper(root, &MAX)
    return MAX
}

func helper(root *TreeNode, MAX *int) int {
    if root == nil {
        return 0
    }
    left, right := helper(root.Left, MAX), helper(root.Right, MAX)
    if left + right + 1 > *MAX {
        *MAX = left + right
    }
    if left > right {
        return left + 1
    } else {
        return right + 1
    }
} 
