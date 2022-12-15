func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    var ans, tmp int
    v1 := prices[0]
    for i := 1; i < len(prices); i++ {
        if prices[i] < v1 {
            v1 = prices[i]
            continue
        }

        tmp = prices[i] - v1
        if tmp > ans {
            ans = tmp
        }
    }

    return ans
}
