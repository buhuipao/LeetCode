// 动态规划
func checkRecord(n int) (ans int) {
    const mod int = 1e9 + 7
    dp := make([][2][3]int, n+1) // 三个维度分别表示：长度，A 的数量，结尾连续 L 的数量
    dp[0][0][0] = 1
    for i := 1; i <= n; i++ {
        // 以 P 结尾的数量
        for j := 0; j <= 1; j++ {
            for k := 0; k <= 2; k++ {
                dp[i][j][0] = (dp[i][j][0] + dp[i-1][j][k]) % mod
            }
        }
        // 以 A 结尾的数量
        for k := 0; k <= 2; k++ {
            dp[i][1][0] = (dp[i][1][0] + dp[i-1][0][k]) % mod
        }
        // 以 L 结尾的数量
        for j := 0; j <= 1; j++ {
            for k := 1; k <= 2; k++ {
                dp[i][j][k] = (dp[i][j][k] + dp[i-1][j][k-1]) % mod
            }
        }
    }
    for j := 0; j <= 1; j++ {
        for k := 0; k <= 2; k++ {
            ans = (ans + dp[n][j][k]) % mod
        }
    }
    return ans
}
