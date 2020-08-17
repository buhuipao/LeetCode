/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    h1, h2 := float64(helper(root.Left)), float64(helper(root.Right))
    return  math.Abs(h1 - h2) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func helper(root *TreeNode) int {
    if root == nil {
        return 0
    }
    L, R := helper(root.Left), helper(root.Right)
    if L > R {
        return L + 1
    } else {
        return R + 1
    }
}
