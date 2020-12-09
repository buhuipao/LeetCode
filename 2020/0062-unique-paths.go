// DP[i][j] = DP[i-1][j] + DP[i][j-1]
/*
func uniquePaths(m int, n int) int {
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, m)
    }
    for i := 0; i < m; i++ {
        dp[0][i] = 1
    }
    for j := 0; j < n; j++ {
        dp[j][0] = 1
    }
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[n-1][m-1]
}
*/

// 优化版DP
func uniquePaths(m int, n int) int {
    dp := make([]int, n)
    for i := range dp {
        dp[i] = 1
    }
    // 一行一行的扫，所以可以压缩成一维的一排：cur[j] = cur[j-1] + old[j]
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[j] = dp[j-1] + dp[j]
        }
    }
    return dp[n-1]
}
