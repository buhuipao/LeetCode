/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    stack := []*TreeNode{root}
    node := root.Left
    ans := []int{}
    for len(stack) != 0 || node != nil {
        for node != nil {
            stack = append(stack, node)
            node = node.Left
        }

        node = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        ans = append(ans, node.Val)

        node = node.Right
    }
    return ans
}
