// dp[i][j] = min(dp[i-1][j]+1, dp[i-1][j-1], dp[i][j-1]+1), when word1[i] == word2[j]
// else: min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
func minDistance(word1 string, word2 string) int {
    n, m := len(word1), len(word2)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, m+1)
    }

    // init
    for i := 0; i <= n; i++ {
        dp[i][0] = i
    }
    for j := 0; j <= m; j++ {
        dp[0][j] = j
    }

    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = min(dp[i-1][j-1], dp[i][j-1]+1, dp[i-1][j]+1)
            } else {
                dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
            }
        }
    }

    return dp[n][m]
}

func min(a, b, c int) int {
    if a > b {
        a = b
    }

    if a > c {
        return c
    }

    return a
}
