/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var ans []int
    stack := []*TreeNode{}
    node := root
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
