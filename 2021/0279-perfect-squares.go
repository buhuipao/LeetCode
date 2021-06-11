// dp
// F(x) = 1 or min(F(j) + F(x-j))，其中 0 < j*j < x
func numSquares(n int) int {
    // 隐含条件：dp[0] = 0
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        // 最坏的情况：n个1
        min := i
        for j := 1; j * j <= i; j++ {
            // dp[i] = MIN(dp[i], dp[i - j * j] + 1)
            // 也可以写成：dp[i] = MIN(dp[i], dp[i - j * j] + dp[j*j])
            if min > dp[i-j*j] + 1 {
                min = dp[i-j*j] + 1
            }
        }
        dp[i] = min
    }
    return dp[n]
}
