/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归后续遍历
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || root == p || root == q {
        return root
    }

    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    // 找到分别找到p、q在root的子树中，于是最近公共祖先就是root
    if left != nil && root != nil {
        return root
    }

    // 在left中找到p或者q
    if left != nil {
        return left
    }

    // 在right中找到p或者q
    if right != nil {
        return right
    }

    // 没有找到
    return None
}
