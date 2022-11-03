/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if p.Val > q.Val {
        return lowestCommonAncestor(root, q, p)
    }

    if root == nil {
        return nil
    }

    if root.Val > p.Val && root.Val > q.Val {
        return lowestCommonAncestor(root.Left, p, q)
    }

    if root.Val < p.Val && root.Val < q.Val {
        return lowestCommonAncestor(root.Right, p, q)
    }

    if root.Val >= p.Val && root.Val <= q.Val {
        return root
    }

    return nil
}
