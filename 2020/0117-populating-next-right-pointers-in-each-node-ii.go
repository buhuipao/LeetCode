/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
// 按层遍历
func connect(root *Node) *Node {
    if root == nil {
        return nil
    }

    last, nlast := root, root
    pre, end := &Node{}, false
    queue := []*Node{root}
    for len(queue) != 0 {
        node := queue[0]
        queue = queue[1:]

        if node.Left != nil {
            queue = append(queue, node.Left)
            nlast = node.Left
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
            nlast = node.Right
        }
        // 不是最后一个节点则连接起来
        if !end {
            pre.Next = node
        }
        // 每行最后一个节点
        if node == last {
            end = true
            last = nlast
        } else {
            end = false
        }
        pre = node
    }
    return root
}
