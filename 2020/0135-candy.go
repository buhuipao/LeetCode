func candy(ratings []int) int {
    n := len(ratings)
    ans, inc, dec, pre := 1, 1, 0, 1
    for i := 1; i < n; i++ {
        // 增长序列
        if ratings[i] >= ratings[i-1] {
            dec = 0
            if ratings[i] == ratings[i-1] { // 如果相同则故意给右边的分配1个糖果以满足要求
                pre = 1
            } else {
                pre++
            }
            ans += pre
            inc = pre   // 记录当前糖果数也就是记录增长序列的步长，因为所有的增长序列都是从1开始的
        } else {
            dec++
            if dec == inc { // 如果递减序列的长度等于上一个增常序列的长度了，
                            // 需要把上一个增长序列的最后一个也开始纳入递减序列，
                            // 避免不满足要求（此处难理解需要画图推导）
                dec++
            }
            ans += dec
            pre = 1
        }
    }
    return ans
}
