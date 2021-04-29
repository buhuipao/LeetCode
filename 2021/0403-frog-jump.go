func canCross(stones []int) bool {
    n := len(stones)
    dp := make([]map[int]bool, n-1)
    for i := range dp {
        dp[i] = map[int]bool{}
    }
    var dfs func(int, int) bool
    dfs = func(i, lastDis int) (res bool) {
        if i == n-1 {
            return true
        }
        if res, has := dp[i][lastDis]; has {
            return res
        }
        defer func() { dp[i][lastDis] = res }()
        for curDis := lastDis - 1; curDis <= lastDis+1; curDis++ {
            if curDis > 0 {
                j := sort.SearchInts(stones, curDis+stones[i])
                if j != n && stones[j] == curDis+stones[i] && dfs(j, curDis) {
                    return true
                }
            }
        }
        return false
    }
    return dfs(0, 0)
}
