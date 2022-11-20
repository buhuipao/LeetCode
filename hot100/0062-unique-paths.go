func uniquePaths(m int, n int) int {
    dp := make([]int, n)
    for i := range dp {
        dp[i] = 1
    }

    // 一行一行的扫，所以可以压缩成一维的一排：cur[j] = cur[j-1] + old[j]
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[j] = dp[j] + dp[j-1]
        }
    }

    return dp[n-1]
}
