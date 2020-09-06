/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
         return nil
    }
    queue := make([]*TreeNode, 1)
    tail, nextTail := root, root
    queue[0] = tail
    ans, tmp := make([][]int, 0), make([]int, 0)
    for len(queue) != 0 {
        node := queue[0]
        queue = queue[1:]
        tmp = append(tmp, node.Val)
        // 更新下一层的尾巴
        if node.Left != nil {
            queue = append(queue, node.Left)
            nextTail = node.Left
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
            nextTail = node.Right
        }
        // 一层到头了
        // 1. 更新tail、2.添加tmp
        if node == tail {
            tail = nextTail
            ans = append(ans, tmp)
            tmp = make([]int, 0)
        }
    }
    l, r := 0, len(ans)-1
    for  l < r {
        ans[l], ans[r] = ans[r], ans[l]
        l++
        r--
    }
    return  ans 
}
