/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    stack := []*TreeNode{root}
    var ans []int
    for len(stack) != 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        ans = append(ans, node.Val)

        if node.Right != nil {
            stack = append(stack, node.Right)
        }

        if node.Left != nil {
            stack = append(stack, node.Left)
        }
    }

    return ans
}
