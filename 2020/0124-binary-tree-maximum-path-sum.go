/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    ans := ^(int(^uint(0)>>1))  // 最小int值
    helper(root, &ans)
    return ans
}

func helper(root *TreeNode, ans *int) int {
    if root == nil {
        return 0
    }
    l := helper(root.Left, ans)
    r := helper(root.Right, ans)
    if l < 0 {
        l = 0
    }
    if r < 0 {
        r = 0
    }
    if root.Val + r + l > *ans {
        *ans = root.Val + r + l
    }
    if l > r {
        return root.Val + l
    } else {
        return root.Val + r
    }
}
