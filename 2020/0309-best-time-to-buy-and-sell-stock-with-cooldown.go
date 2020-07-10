// DP
func maxProfit(prices []int) int {
    n := len(prices)
    if n < 2 {
        return 0
    }
    a, b := make([]int, 3), make([]int, 3)
    // dp[i][0]：本来就不持有股票，1）要么是由于冷冻期没法买；2）至少2天前就卖了，到现在还是不持有
    // dp[i][1]：持有股票，1）前一天就持有；2）今天购买了
    // dp[i][2]：今天卖出导致不持有股票，1）今天卖出股票
    a = []int{0, 0-prices[0], 0}
    for i := 1; i < n; i++ {
        a, b[0], b[1], b[2] = b, max(a[2], a[0]), max(a[1], a[0]-prices[i]), a[1] + prices[i]
    }
    return max(b[0], b[2])
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
