func maxProfit(k int, prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    // 初始化第0天的数据
    // 购买收益b中，应该都为minInt
    // 卖出收益，应该为0
    b, s := make([]int, k+1), make([]int, k+1)
    for i := range b {
        b[i] = ^int(^uint(0)>>1)
    }

    for _, p := range prices {
        for i := 1; i <= k; i++ {
            b[i] = max(b[i], s[i-1]-p)
            s[i] = max(s[i], b[i]+p)
        }
    }

    return s[k]
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
