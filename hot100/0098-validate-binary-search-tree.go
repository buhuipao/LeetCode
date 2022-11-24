/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    _, _, ok := helper(root)
    return ok
}

func helper(root *TreeNode) (int, int, bool) {
    if root == nil {
        return 0, 0, true
    }

    if root.Left == nil && root.Right == nil {
        return root.Val, root.Val, true
    }

    maxV, minV := root.Val, root.Val
    if root.Left != nil {
        maxL, minL, okL := helper(root.Left)
        if !okL {
            return 0, 0, false
        }

        if root.Val <= maxL {
            return 0, 0, false
        }

        minV = minL
    }

    if root.Right != nil {
        maxR, minR, okR := helper(root.Right)
        if !okR {
            return 0, 0, false
        }

        if root.Val >= minR {
            return 0, 0, false
        }

        maxV = maxR
    }

    return maxV, minV, true
}
