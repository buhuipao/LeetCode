/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 树形前缀和
func pathSum(root *TreeNode, targetSum int) int {
    if root == nil {
        return 0
    }

    var ans int
    m := map[int]int{
        0: 1,
    }
    dfs(root, &ans, m, targetSum, 0)

    return ans
}

func dfs(root *TreeNode, ans *int, m map[int]int, targetSum, cur int) {
    if root == nil {
        return
    }
    
    cur += root.Val
    *ans += m[cur-targetSum]
    m[cur]++
    dfs(root.Left, ans, m, targetSum, cur)
    dfs(root.Right, ans, m, targetSum, cur)
    m[cur]--
}
