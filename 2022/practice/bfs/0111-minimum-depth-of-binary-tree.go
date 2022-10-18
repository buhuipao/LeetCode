/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    var ans int
    nodes := []*TreeNode{root}
    for len(nodes) != 0 {
        n := len(nodes)
        for i := 0; i < n; i++ {
            if nodes[i].Left == nil && nodes[i].Right == nil {
                return ans+1
            }

            if nodes[i].Left != nil {
                nodes = append(nodes, nodes[i].Left)
            }

            if nodes[i].Right != nil {
                nodes = append(nodes, nodes[i].Right)
            }
        }
        ans++
        nodes = nodes[n:]
    }

    return ans
}
