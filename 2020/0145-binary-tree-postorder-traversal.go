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
    s1, s2 := []*TreeNode{}, []*TreeNode{}
    s1 = append(s1, root)
    for len(s1) != 0 {
        node := s1[len(s1)-1]
        s1 = s1[:len(s1)-1]
        s2 = append(s2, node)

        // 由于用到了两个栈，于是乎 压栈时顺利和遍历顺序一致
        if node.Left != nil {
            s1 = append(s1, node.Left)
        }
        if node.Right != nil {
            s1 = append(s1, node.Right)
        }
    }
    ans := make([]int, 0, len(s2))
    for i := len(s2)-1; i >= 0; i-- {
        ans = append(ans, s2[i].Val)
    }
    return ans
}
