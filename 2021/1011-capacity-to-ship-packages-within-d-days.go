func shipWithinDays(weights []int, D int) int {
    // 确定二分查找左右边界
    left, right := 0, 0
    for _, w := range weights {
        if w > left {
            left = w
        }
        right += w
    }
    return left + sort.Search(right-left, func(x int) bool {
        x += left
        day := 1 // 需要运送的天数
        sum := 0 // 当前这一天已经运送的包裹重量之和
        for _, w := range weights {
            if sum+w > x {
                day++
                sum = 0
            }
            sum += w
        }
        return day <= D
    })
}
