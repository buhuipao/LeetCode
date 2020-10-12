/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    if root == nil {
        return 0
    }
    stack := []*TreeNode{root}
    node := root.Left
    ans, pre := int(^uint(0)>>1), -1
    for len(stack) != 0 || node != nil {
        for node != nil {
            stack = append(stack, node)
            node = node.Left
        }
        node = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if pre != -1 && node.Val - pre < ans {
            ans = node.Val - pre
        }
        node, pre = node.Right, node.Val
    }
    return ans
}
