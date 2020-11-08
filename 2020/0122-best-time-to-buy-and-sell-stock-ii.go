// 累加单调递增的值即可
func maxProfit(prices []int) int {
    pre := int(^uint(0) >> 1)
    ans := 0
    n := len(prices)
    for i := 0; i < n; i++ {
        if prices[i] > pre {
            ans += (prices[i] - pre)
        }
        pre = prices[i]
    }

    return ans
}
