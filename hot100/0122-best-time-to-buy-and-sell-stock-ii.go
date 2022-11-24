// 找出所有单调递增区间
func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    var ans int
    for i := 1; i < len(prices); i++ {
        ans += max(0, prices[i]-prices[i-1])
    }
    
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
