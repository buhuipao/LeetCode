/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
    var ans int

    sum(root, &ans)

    return ans
}

func sum(root *TreeNode, ans *int) int {
    if root == nil {
        return 0
    }

    L := sum(root.Left, ans)
    R := sum(root.Right, ans)

    if L > R {
        *ans += L - R
    } else {
        *ans += R - L
    }

    return L + R + root.Val
}
