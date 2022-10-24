/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
    rob, notRob := robRoot(root)
    return max(rob, notRob)
}

func robRoot(root *TreeNode) (robCur int, notRobCur int) {
    if root == nil {
        return
    }

    robLeft, notRobLeft := robRoot(root.Left)
    robRight, notRobRight := robRoot(root.Right)
    robCur = root.Val + notRobLeft + notRobRight
    notRobCur = max(robLeft, notRobLeft) + max(robRight, notRobRight)

    return robCur, notRobCur
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
