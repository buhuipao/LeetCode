// Hash表
// 关键是找到nums[0]，只要找到它就可以推出整个序列
func restoreArray(adjacentPairs [][]int) []int {
    n := len(adjacentPairs) + 1
    g := make(map[int][]int, n)
    for _, p := range adjacentPairs {
        v, w := p[0], p[1]
        g[v] = append(g[v], w)
        g[w] = append(g[w], v)
    }

    ans := make([]int, n)
    for e, adj := range g {
        if len(adj) == 1 {
            ans[0] = e
            break
        }
    }

    ans[1] = g[ans[0]][0]
    for i := 2; i < n; i++ {
        adj := g[ans[i-1]]
        if ans[i-2] == adj[0] {
            ans[i] = adj[1]
        } else {
            ans[i] = adj[0]
        }
    }
    return ans
}
