/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
}
*/

func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    var ans []int
    var stack []*TreeNode
    node := root
    for len(stack) != 0 || node != nil {
        for node != nil {
            // 先把自己放进去
            stack = append(stack, node)
            // 再把自己的左儿子放进去，这样下方弹栈时可以确保左儿子最先出栈、然后是自己
            node = node.Left
        }

        node, stack = stack[len(stack)-1], stack[:len(stack)-1]
        ans = append(ans, node.Val)

        // 这里尝试取到自己的右儿子，回到上方循环后可以压入栈
        node = node.Right
    }

    return ans
}
