/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }
    var a, b, c int
    a = helper(root, sum)
    b = pathSum(root.Left, sum)
    c = pathSum(root.Right, sum)
    return a + b + c
}

// 以root节点为起始位置和为target 的数量
func helper(root *TreeNode, target int) int {
    if root == nil {
        return 0
    }
    var c int
    if root.Val == target {
        c = 1
    }
    return c + helper(root.Left, target-root.Val) + helper(root.Right, target-root.Val)
}
