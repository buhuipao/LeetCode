var minInt = ^int(^uint(0) >> 1)

func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    // 初始化股票买卖情况
    d0 := 0             // 没有任何买入操作
    d1 := -prices[0]    // 买入一支股票
    d2 := 0             // 买入一支，卖出一支
    d3 := -prices[0]    // 买入两支，卖出一支
    d4 := 0             // 买入两支，卖出两支

    for i := 1; i < len(prices); i++ {
        d1 = max(d1, d0-prices[i])
        d2 = max(d2, d1+prices[i])
        d3 = max(d3, d2-prices[i])
        d4 = max(d4, d3+prices[i])
    }

    return max(d0, d1, d2, d3, d4)
}

func max(nums... int) int {
    if len(nums) == 0 {
        return 0
    }

    n := nums[0]
    for i := range nums {
        if n < nums[i] {
            n = nums[i]
        }
    }

    return n
}
