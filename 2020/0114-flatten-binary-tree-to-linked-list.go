/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // 先序遍历
func flatten(root *TreeNode)  {
    if root == nil {
        return
    }
    stack := []*TreeNode{root}
    pre := &TreeNode{}
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
        // 清理
        node.Left, node.Right = nil, nil
        pre, pre.Right = node, node
    }
}
