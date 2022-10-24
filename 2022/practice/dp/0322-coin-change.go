func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i := range dp {
        // 故意设置为极大值
        dp[i] = amount+1
    }
    dp[0] = 0

    for i := 0; i <= amount; i++ {
        for j := range coins {
            if i - coins[j] < 0 {
                continue
            }

            if dp[i-coins[j]]+1 < dp[i] {
                dp[i] = dp[i-coins[j]] + 1
            } 
        }
    }

    if dp[amount] > amount {
        return -1
    }

    return dp[amount]
}
