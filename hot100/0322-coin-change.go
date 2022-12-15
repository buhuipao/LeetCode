func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i := range dp {
        // 故意设置为极大
        dp[i] = amount + 1
    }
    dp[0] = 0

    for i := 1; i <= amount; i++ {
        for j := range coins {
            if i - coins[j] < 0 {
                continue
            }

            dp[i] = min(dp[i], dp[i-coins[j]]+1)
        }
    }

    // 证明没有合适的方案
    if dp[amount] > amount {
        return -1
    }

    return dp[amount]
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
