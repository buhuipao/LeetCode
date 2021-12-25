/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isEvenOddTree(root *TreeNode) bool {
    if root == nil {
        return true
    }

    var queue []*TreeNode
    var pre, level int
    queue = append(queue, root)
    nextR := root
    for len(queue) != 0 {
        node := queue[0]
        queue = queue[1:]

        if node.Left != nil {
            queue = append(queue, node.Left)
        }

        if node.Right != nil {
            queue = append(queue, node.Right)
        }

        if level % 2 == 0 {
            if node.Val % 2 != 1 {
                return false
            }

            if pre != 0 && pre >= node.Val {
                return false
            }
        } else {
            if node.Val % 2 != 0 {
                return false
            }

            if pre != 0 && pre <= node.Val {
                return false
            }
        }

        pre = node.Val

        if node == nextR {  // 换行
            level++
            if len(queue) > 0 {
                nextR = queue[len(queue)-1]
            }

            pre = 0
        }
    }

    return true
}
