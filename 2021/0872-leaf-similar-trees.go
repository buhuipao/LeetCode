/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // 中序遍历
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    l1, l2 := visit(root1), visit(root2)
    if len(l1) != len(l2) {
        return false
    }
    for i := range l1 {
        if l1[i] != l2[i] {
            return false
        }
    }
    return true
}

func visit(root *TreeNode) []int {
    stack := make([]*TreeNode, 0)
    node := root
    ans := make([]int, 0)
    for len(stack) != 0 || node != nil {
        for node != nil {
            stack = append(stack, node)
            node = node.Left
        }
        node = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if node.Left == nil && node.Right == nil {
            ans = append(ans, node.Val)
        }
        node = node.Right
    }
    return ans
}
