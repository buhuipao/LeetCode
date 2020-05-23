/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
 // 反向中序遍历
func convertBST(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    sum := 0
    stack := []*TreeNode{root}
    node := root.Right
    for len(stack) != 0 || node != nil {
        for node != nil {
            stack = append(stack, node)
            node = node.Right
        }
        node = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        sum += node.Val
        node.Val = sum
        node = node.Left
    }

    return root
}
