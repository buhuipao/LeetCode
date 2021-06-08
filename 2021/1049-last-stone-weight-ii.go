// 背包问题
// 需要找到能凑出的最大的不超过sum(nums)/2的值X，最后剩下的就是：sum(nums) - 2 *X
func lastStoneWeightII(stones []int) int {
    sum := 0
    for _, v := range stones {
        sum += v
    }
    n, m := len(stones), sum/2
    dp := make([][]bool, n+1)
    for i := range dp {
        dp[i] = make([]bool, m+1)
    }
    dp[0][0] = true
    for i, weight := range stones {
        for j := 0; j <= m; j++ {
            if j < weight {
                dp[i+1][j] = dp[i][j]
            } else {
                dp[i+1][j] = dp[i][j] || dp[i][j-weight]
            }
        }
    }
    for j := m; ; j-- {
        if dp[n][j] {
            return sum - 2*j
        }
    }
}
