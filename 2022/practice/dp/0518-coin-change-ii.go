func change(amount int, coins []int) int {
    dp := make([]int, amount+1)
    dp[0] = 1

    for j := range coins {
        for i := 1; i <= amount; i++ {
            if i - coins[j] >= 0 {
                dp[i] += dp[i-coins[j]]
            }
        }
    }

    return dp[amount]
}
