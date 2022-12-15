/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }

    l := invertTree(root.Left)
    r := invertTree(root.Right)

    root.Left, root.Right = r, l

   return root
}
