/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    if root == nil {
        return nil
    }
    last, nlast := root, root
    queue := []*TreeNode{root}
    ans := make([]float64, 0)
    sum, count := float64(0), 0
    for len(queue) != 0 {
        node := queue[0]
        sum += float64(node.Val)
        count++
        queue = queue[1:]
        if node.Left != nil {
            queue  = append(queue, node.Left)
            nlast = node.Left
        }
        if node.Right != nil {
            queue  = append(queue, node.Right)
            nlast = node.Right
        }
        if node == last {
            last = nlast
            ans = append(ans, sum/float64(count))
            sum, count = float64(0), 0
        }
    }
    return ans
}
