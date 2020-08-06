/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil || q == nil {
        return p == q
    }
    if q.Val != p.Val {
        return false
    }
    return isSameTree(q.Left, p.Left) && isSameTree(q.Right, p.Right)
}
