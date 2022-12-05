/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    ans := ^int(^uint(0)>>1) // 最小值
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

    // 更新ans
    tmp := root.Val + l + r
    if tmp > *ans {
        *ans = tmp
    }

    // 选择更大值的路径返回给上层，因为路径只能选择：root、root+root.Left、root+root.Right，不能选择整个子树
    if l > r {
        return root.Val + l
    } else {
        return root.Val + r
    }
}
