/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }

    var hl, hr int
    left, right := root, root
    for left != nil {
        left = left.Left
        hl++
    }
    for right != nil {
        right = right.Right
        hr++
    }

    // 满二叉树
    if hl == hr {
        return int(math.Pow(float64(2), float64(hl)) - 1)
   }

   return 1 + countNodes(root.Left) + countNodes(root.Right)
}
