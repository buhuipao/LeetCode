
func numTrees(n int) int {
    dp := make(map[int]int, n+1)
    // 空树也是一种二叉树
    dp[0], dp[1] = 1, 1
    for i := 2; i <=n; i++ {
        for j := 0; j < i; j++ {
            dp[i] += dp[j] * dp[i-j-1]
        }
    }
    return dp[n]
}
