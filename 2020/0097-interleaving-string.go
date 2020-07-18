// DP
func isInterleave(s1 string, s2 string, s3 string) bool {
    // 特判
    if len(s3) != len(s1) + len(s2) {
        return false
    }
    n1, n2 := len(s1), len(s2)
    dp := make([][]bool, n1+1)
    for i := range dp {
        dp[i] = make([]bool, n2+1)
    }
    dp[0][0] = true
    for i := 1; i <= n1; i++ {
        if s1[i-1] == s3[i-1] {
            dp[i][0] = true
        } else {
            break
        }   
    }
    for j := 1; j <= n2; j++ {
        if s2[j-1] == s3[j-1] {
            dp[0][j] = true
        } else {
            break
        }
    }
    // 建立一个表格
    for i := 1; i <= n1; i++ {
        for j := 1; j <= n2; j++ {
            dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
        }
    }
    return dp[n1][n2]
}
