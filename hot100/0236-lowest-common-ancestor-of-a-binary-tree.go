/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    // 找到一个节点
    if root == nil || root == p || root == q {
        return root
    }

    left := lowestCommonAncestor(root.Left, p, q)
    rihght := lowestCommonAncestor(root.Right, p, q)
    // 找到分别找到p、q在root的子树中，于是最近公共祖先就是root
    if left != nil && rihght != nil {
        return root
    }

    // 在left中找到p或者q
    if left != nil {
        return left
    }

    // 在right中找到p或者q
    if rihght != nil {
        return rihght
    }

    return nil
}
