// dp
func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    // 第1天结束时的状态，{不持有股票，持有股票}
    dp := [2]int{0, 0-prices[0]}
    for i := 1; i < len(prices); i++ {
        // 第i天结束时，
        // 不持有股票(dp[0])：max(今天不买--继承昨天的无股票收益，今天卖了股票--昨天手里持有股票收益+今天卖出股票价格)
        // 持有股票(dp[1])：max(今天不卖--继承昨天的收益和持有状态，今天买了股票--昨天手里无股票收益-今天买入股票价格)
        dp[0], dp[1] = max(dp[0], dp[1]+prices[i]), max(dp[1], dp[0]-prices[i]) 
    }

    return max(dp[0], dp[1])
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
