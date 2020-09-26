/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
    ans := make([][]int, 0)

    var h func(*TreeNode, int, []int)
    h = func(root *TreeNode, sum int, path []int) {
        if root == nil {
            return
        }
        // 符合条件则加入最终结果
        if root.Left == nil && root.Right == nil && sum == root.Val {
            newP := make([]int, len(path))
            copy(newP, path)
            ans = append(ans, append(newP, root.Val))
            return
        }
        // 尝试左右子树
        path = append(path, root.Val)
        h(root.Left, sum-root.Val, path)
        h(root.Right, sum-root.Val, path)
    }
    
    h(root, sum, []int{})

    return ans
}
