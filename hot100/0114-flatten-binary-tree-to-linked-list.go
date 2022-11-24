/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    if root == nil {
        return
    }

    l, r := root.Left, root.Right
    flatten(root.Left)
    flatten(root.Right)

    // 根据先序遍历展开的要求，把左子树放到链表当前节点的后一个节点（也就是root.Right）
    root.Left, root.Right = nil, l
    // 经过上面的 flatten(root.Left) 操作，可以知道：root的左子树已经全部展开为root右侧的后继节点了
    // 于是需要持续查找右侧后继节点，并将原先的右子树放到最后一个后继节点后
    for root.Right != nil {
        root = root.Right
    }

    root.Right = r
}
