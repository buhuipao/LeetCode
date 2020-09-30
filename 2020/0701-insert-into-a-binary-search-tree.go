/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 中序遍历
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }
    stack := []*TreeNode{root}
    node := root.Left
    var pre *TreeNode
    for len(stack) != 0 || node != nil {
        for node != nil {
            stack = append(stack, node)
            node = node.Left
        }
        node = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        // 找到位置
        if node.Val >= val {
            node.Val, val = val, node.Val
        }

        node, pre = node.Right, node
    }
    // 转移最后一个节点
    pre.Right = &TreeNode{Val: val}

    return root
}
