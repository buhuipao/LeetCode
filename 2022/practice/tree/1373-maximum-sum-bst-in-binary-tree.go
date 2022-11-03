/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxSumBST(root *TreeNode) int {
    var ans int
    visit(root, &ans)
    return ans
}

func visit(root *TreeNode, ans *int) (int, int, int, bool) {
    if root == nil {
        return 0, 0, 0, true
    }

    l, minL, maxL, ok1 := visit(root.Left, ans)
    r, minR, maxR, ok2 := visit(root.Right, ans)
    // 左右子树是否为BST
    if !ok1 || !ok2 {
        return 0, 0, 0, false
    }

    sum, minV, maxV := root.Val, root.Val, root.Val
    // 是否大于左子树最大值
    if root.Left != nil {
        if !(root.Val > maxL) {
            return 0, 0, 0, false
        }
        minV = minL
        sum += l
    }

    // 是否小于右子树最小值
    if root.Right != nil {
        if !(root.Val < minR) {
            return 0, 0, 0, false
        }
        maxV = maxR
        sum += r
    }

    *ans = max(*ans, sum)

    return sum, minV, maxV, true
}

func min(a, b int) int{
    if a > b {
        return b
    }

    return a
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
