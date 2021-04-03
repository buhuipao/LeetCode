// dp[i][j] = max(dp[i][j-1], dp[i-1][j]) or dp[i-1][j-1] + 1
func longestCommonSubsequence(text1 string, text2 string) int {
    m, n := len(text1), len(text2)
    // 初始化
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    // 计算
    for i, c1 := range text1 {
        for j, c2 := range text2 {
            if c1 == c2 {
                dp[i+1][j+1] = dp[i][j] + 1
            } else {
                dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
            }
        }
    }
    return dp[m][n]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
