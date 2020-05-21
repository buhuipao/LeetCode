/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    n := len(inorder)
    // 特判
    if n == 0 {
        return nil
    }
    // 找到所有左子树的节点
    var i int
    for i = 0; i < n; i++ {
        if inorder[i] == preorder[0] {
            break
        }
    }
    left := buildTree(preorder[1:i+1], inorder[0:i])
    right := buildTree(preorder[i+1:], inorder[i+1:])

    return &TreeNode{Val: preorder[0], Left: left, Right: right}
}
