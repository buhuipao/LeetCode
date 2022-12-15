// F(x) = 1 or min(F(j) + F(x-j))，其中 0 < j*j < x
func numSquares(n int) int {
    // 隐含条件：dp[0] = 0
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dp[i] = i
        for j := 1; j*j <= i; j++ {
            dp[i] = min(dp[i], dp[i-j*j]+1)
        }
    }

    return dp[n]
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
