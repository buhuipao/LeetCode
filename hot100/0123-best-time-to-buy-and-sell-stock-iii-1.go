func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    // 初始化
    s, b := make([]int, 3), make([]int, 3)
    for i := range b {
        b[i] = -prices[0]
    }

    for i := range prices {
        for j := 1; j < len(s); j++ {
            b[j], s[j] = max(b[j], s[j-1]-prices[i]), max(s[j], b[j]+prices[i])
        }
    }

    return s[2]
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
