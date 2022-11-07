/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    var s1, s2 []*TreeNode
    var ans []int
    s1 = append(s1, root)
    for len(s1) != 0 {
        node := s1[len(s1)-1]
        s2 = append(s2, node)
        s1 = s1[:len(s1)-1]

        if node.Left != nil {
            s1 = append(s1, node.Left)
        }

        if node.Right != nil {
            s1 = append(s1, node.Right)
        }
    } 

    for i := len(s2)-1; i >= 0; i-- {
        ans = append(ans, s2[i].Val)
    }

    return ans
}
