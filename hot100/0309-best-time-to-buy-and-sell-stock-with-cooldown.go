func maxProfit(prices []int) int {
    n := len(prices)
    if n < 2 {
        return 0
    }

    // 三种状态：
    // a：不持有股票，1）要么是由于冷冻期没法买；2）至少2天前就卖了，到现在还是不持有
    // b：持有股票，1）前一天就持有；2）今天新购买了
    // c：刚卖出股票：导致不持有股票，1）今天卖出股票

    // 为什么会比「无冷冻期」多出一种「刚卖出」状态？
    // 答：因为在无冷冻期时，『持有股票』可以简单的从『不持有股票状态』通过购买股票转换，因而不能把『刚卖出股票』和 『不持有股票--冷冻期』混在一起
    a, b, c := 0, ^(int(^uint(0)>>1)), 0
    for i := range prices {
        // a = max(c, a) 中max函数的第二参数a，就是昨天不持有股票的状态，所以是直接继承之前的a
        a, b, c = max(c, a), max(b, a-prices[i]), b+prices[i]
    }

    return max(a, c)
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
