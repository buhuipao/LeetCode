/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    var ans int

    helper(root, &ans)

    return ans
}

func helper(root *TreeNode, ans *int) int {
    if root == nil {
        return 0
    }

    l, r := helper(root.Left, ans), helper(root.Right, ans)
    *ans = max(l+r, *ans)

    return 1 + max(l, r)
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
