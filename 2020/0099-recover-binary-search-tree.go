/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // 中序遍历
func recoverTree(root *TreeNode)  {
    if root == nil {
        return
    }
    // 两个待交换的节点
    var node1, node2 *TreeNode
    stack := []*TreeNode{root}
    node := root.Left
    // 设定pre为最小值
    MIN_INT := ^(int(^uint(0)>>1))
    pre := &TreeNode{Val: MIN_INT}
    for len(stack) != 0 || node != nil {
        for node != nil {
            stack = append(stack, node)
            node = node.Left
        }
        node = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        // 找到第一个错误节点
        if node1 == nil && pre.Val > node.Val {
            node1 = pre
        }
        // 找到第二错误节点
        if node1 != nil && pre.Val > node.Val {
            node2 = node
        }
        // 重新赋值
        node, pre = node.Right, node
    }
    // 调换错误节点的值
    if node1 != nil {
        node1.Val, node2.Val = node2.Val, node1.Val
    }
}
