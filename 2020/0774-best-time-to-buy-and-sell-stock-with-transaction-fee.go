func maxProfit(prices []int, fee int) int {
    n := len(prices)
    dp := make([][2]int, n)
    // dp[0] = [0,-prices[0]]
    dp[0][1] = -prices[0]
    for i := 1; i < n; i++ {
        // 今天不持有股票的两种情况：1）昨天也没有；2）昨天但有今天卖了，需要扣除手续费
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
        // 今天持有股票的两种情况：1）昨天就有；2）昨天没有但今天卖出了，暂时不算入手续费
        dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
    }
    return dp[n-1][0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
