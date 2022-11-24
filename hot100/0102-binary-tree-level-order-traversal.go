/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    last, nlast := root, root
    var ans [][]int
    var tmp []int
    queue := []*TreeNode{root}
    for len(queue) != 0 {
        node := queue[0]
        queue = queue[1:]
        tmp = append(tmp, node.Val)
        if node.Left != nil {
            queue = append(queue, node.Left)
            nlast = node.Left
        }

        if node.Right != nil {
            queue = append(queue, node.Right)
            nlast = node.Right
        }

        if last == node {
            last = nlast
            ans = append(ans, tmp)
            tmp = []int{}
        }
    }

    return ans
}
