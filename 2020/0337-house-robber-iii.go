/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // 递归
func rob(root *TreeNode) int {
    ret := helper(root)
    return max(ret)
}

// 返回偷/不偷当前节点的结果
func helper(root *TreeNode) []int {
    if root == nil {
        return []int{0, 0}
    }
    L, R := helper(root.Left), helper(root.Right)

    // 偷当前节点，所以就不能偷两个孩子
    a := root.Val + L[1] + R[1]
    // 不偷当前节点, 所以两个孩子可偷可不偷，取最大即可
    b := 0 + max(L) + max(R)

    return []int{a, b}
}

func max(a []int) int {
    if a[0] > a[1] {
        return a[0]
    }
    return a[1]
}

