/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    ans, tmp := make([][]int, 0), make([]int, 0)
    last, nlast := root, root
    count := 0
    for len(queue) != 0 {
        // 出队
        node := queue[0]
        queue = queue[1:]
        tmp = append(tmp, node.Val)
        // 更新下一行的最后一个
        if node.Left != nil {
            queue = append(queue, node.Left)
            nlast = node.Left
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
            nlast = node.Right
        }
        // 换行
        if node == last {
            if count % 2 != 0 {
                m := len(tmp)
                for i := 0; i < m/2; i++ {
                    tmp[i], tmp[m-1-i] = tmp[m-1-i], tmp[i]
                } 
            }
            count++
            ans = append(ans, tmp)
            tmp = tmp[len(tmp):]
            last = nlast
        }
    }
    return ans
}
