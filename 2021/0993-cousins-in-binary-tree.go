/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isCousins(root *TreeNode, x int, y int) bool {
    ret1, ret2 := findNode(&TreeNode{}, root, x, 0), findNode(&TreeNode{}, root, y, 0)
    return ret1[0] == ret2[0] && ret1[1] != ret2[1]
}

// 查找节点，返回值：[层级，父节点值]
func findNode(father, root *TreeNode, v, k int) []int {
    if root == nil {
        return nil
    }
    if root.Val == v {
        return []int{k, father.Val}
    }
    L, R := findNode(root, root.Left, v, k+1), findNode(root, root.Right, v, k+1)
    if len(L) != 0 {
        return L
    }
    return R
}
