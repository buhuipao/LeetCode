// DFS
func numWays(n int, relation [][]int, k int) (ans int) {
    edges := make([][]int, n)
    for _, r := range relation {
        src, dst := r[0], r[1]
        edges[src] = append(edges[src], dst)
    }

    var dfs func(int, int)
    dfs = func(x, step int) {
        if step == k {
            if x == n-1 {
                ans++
            }
            return
        }
        for _, y := range edges[x] {
            dfs(y, step+1)
        }
    }
    
    dfs(0, 0)

    return
}
