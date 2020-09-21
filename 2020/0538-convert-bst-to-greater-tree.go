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
    stack := []*TreeNode{root}
    node := root.Right
    var sum int
    for len(stack) != 0 || node != nil {
        for node != nil {
            stack = append(stack, node)
           node = node.Right
        }

        node = stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        node.Val += sum
        sum = node.Val

        node = node.Left
    }

    return root
}
