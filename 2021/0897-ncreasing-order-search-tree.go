/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func increasingBST(root *TreeNode) *TreeNode {
    stack, node := []*TreeNode{}, root
    var pre, newRoot *TreeNode
    for len(stack) != 0 || node != nil {
        for node != nil {
            stack = append(stack, node)
            node = node.Left
        }
        node = stack[len(stack)-1]
        node.Left = nil
        stack = stack[:len(stack)-1]
        if newRoot == nil {
            newRoot = node
        }
        if pre != nil {
            pre.Right = node
        }
        node, pre = node.Right, node
    }
    return newRoot
}
