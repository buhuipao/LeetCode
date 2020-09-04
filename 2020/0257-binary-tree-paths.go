/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func binaryTreePaths(root *TreeNode) []string {
    ans := helper(root)
    return ans
}

func helper(root *TreeNode) []string {
    if root == nil {
        return nil
    }
    if root.Right == nil && root.Left == nil {
        return []string{fmt.Sprintf("%d", root.Val)}
    }
    var L, R, ans []string
    if root.Left != nil {
        L = helper(root.Left)
        for i := range L {
            ans = append(ans, fmt.Sprintf("%d->%s", root.Val, L[i]))
        }
    }
    if root.Right != nil {
        R = helper(root.Right)
        for i := range R {
            ans = append(ans, fmt.Sprintf("%d->%s", root.Val, R[i]))
        }
    }
    return ans 
}
