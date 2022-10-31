/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    type pair struct {
        node *TreeNode
        idx  int
    }
    repeat := map[*TreeNode]struct{}{}
    seen := map[[3]int]pair{}
    // 自编号，单调递增
    // 用于给第一次遇到的节点编号
    idx := 0

    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        // 基础编号为0
        if node == nil {
            return 0
        }
        
        tri := [3]int{node.Val, dfs(node.Left), dfs(node.Right)}
        // 如果发现两个子树和root.Val都相同，证明之前遇到过，
        // 那么直接返回之前节点的编号，以便进行更上层对比是否相同
        if p, ok := seen[tri]; ok {
            repeat[p.node] = struct{}{}
            return p.idx
        }
        // 自编号自增
        idx++
        seen[tri] = pair{node, idx}
        return idx
    }
    dfs(root)

    ans := make([]*TreeNode, 0, len(repeat))
    for node := range repeat {
        ans = append(ans, node)
    }

    return ans
}
