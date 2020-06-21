/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    var ans int
    // MIN INT
    ans = ^int(^uint(0)>>1)
    helper(root, &ans)
    return ans
}

func helper(root *TreeNode, ans *int) int {
    if root == nil {
        return 0
    }
    l := helper(root.Left, ans)
    if l < 0 {
        l = 0
    }
    r := helper(root.Right, ans)
    if r < 0 {
        r = 0
    }
    if root.Val + l + r > *ans {
        *ans = root.Val + l + r
    }
    if l > r {
        return root.Val + l
    } else {
        return root.Val + r
    }
}
