/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    stack := []*TreeNode{}
    node :=root
    var ans int
    for node != nil || len(stack) != 0 {
        for node != nil {
            stack = append(stack, node)
            node = node.Left
        }

        node = stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        k--
        if k == 0 {
            return node.Val
        }
        
        node = node.Right
    }

    return ans
}
