// COPY FROM LeetCode-Solution
func numWays(steps, arrLen int) int {
    const mod = 1e9 + 7
    maxColumn := min(arrLen-1, steps)
    dp := make([][]int, steps+1)
    for i := range dp {
        dp[i] = make([]int, maxColumn+1)
    }
    dp[0][0] = 1
    for i := 1; i <= steps; i++ {
        for j := 0; j <= maxColumn; j++ {
            dp[i][j] = dp[i-1][j]
            if j-1 >= 0 {
                dp[i][j] = (dp[i][j] + dp[i-1][j-1]) % mod
            }
            if j+1 <= maxColumn {
                dp[i][j] = (dp[i][j] + dp[i-1][j+1]) % mod
            }
        }
    }
    return dp[steps][0]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
