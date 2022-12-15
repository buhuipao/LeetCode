/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
    robRoot, notRobRoot := robNode(root)

    return max(robRoot, notRobRoot)
}

func robNode(root *TreeNode) (int, int) {
    if root == nil {
        return 0, 0
    }

    robLeft, notRobLeft := robNode(root.Left)
    robRight, notRobRight := robNode(root.Right)

    return notRobLeft+notRobRight+root.Val, max(robLeft, notRobLeft)+max(robRight, notRobRight)
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
