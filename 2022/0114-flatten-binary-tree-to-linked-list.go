/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归解法
func flatten(root *TreeNode)  {
    if root == nil {
        return
    }

    left, right := root.Left, root.Right
    flatten(root.Left)
    flatten(root.Right)

    root.Left = nil
    root.Right = left

    // 先找到左子树的最右边儿子，然后将右子树连接至左之树最右边
    for root.Right != nil {
        root = root.Right
    }

    root.Right = right
}
