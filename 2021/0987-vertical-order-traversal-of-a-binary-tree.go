/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // 暴力解法
type data struct{ col, row, val int }

func verticalTraversal(root *TreeNode) (ans [][]int) {
    nodes := []data{}
    var dfs func(*TreeNode, int, int)
    dfs = func(node *TreeNode, row, col int) {
        if node == nil {
            return
        }
        nodes = append(nodes, data{col, row, node.Val})
        dfs(node.Left, row+1, col-1)
        dfs(node.Right, row+1, col+1)
    }
    dfs(root, 0, 0)

    sort.Slice(nodes, func(i, j int) bool {
        a, b := nodes[i], nodes[j]
        return a.col < b.col || a.col == b.col && (a.row < b.row || a.row == b.row && a.val < b.val)
    })

    lastCol := math.MinInt32
    for _, node := range nodes {
        if node.col != lastCol {
            lastCol = node.col
            ans = append(ans, nil)
        }
        ans[len(ans)-1] = append(ans[len(ans)-1], node.val)
    }
    return
}
