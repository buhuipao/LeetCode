// 计数法
// 对'*'进行猜测, 然后可以统计左括号的最小计数。
// 最小计数为0则证明是有效的。
// 反问：为什么不能是最大计数是0，因为最大计数为0，则证明过程中早就出现了计数为负的情况，直接返回false了
func checkValidString(s string) bool {
    var minCount, maxCount int
    for _, ch := range s {
        if ch == '(' {
            minCount++
            maxCount++
        } else if ch == ')' {
            minCount = max(minCount-1, 0)
            maxCount--
            if maxCount < 0 {
                return false
            }
        } else {
            minCount = max(minCount-1, 0)
            maxCount++
        }
    }

    return minCount == 0
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}
