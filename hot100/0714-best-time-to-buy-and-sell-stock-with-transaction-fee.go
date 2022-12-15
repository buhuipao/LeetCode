func maxProfit(prices []int, fee int) int {
    if len(prices) < 2 {
        return 0
    }

    // 初始状态
    // Y：持有股票, N：不持有股票
    Y, N := -5*10*10*10*10, 0 
    for i := range prices {
        Y, N = max(Y, N-prices[i]), max(N, Y+prices[i]-fee)
   }

    return max(Y, N)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}
