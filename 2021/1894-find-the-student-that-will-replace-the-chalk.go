// 前缀和
func chalkReplacer(chalk []int, k int) int {
    var sum int
    for i := range chalk {
        chalk[i] += sum
        sum = chalk[i]
    }

    k = k % sum
    l, r := 0, len(chalk)-1
    for l <= r {
        m := l + (r - l) / 2
        if chalk[m] <= k {
            l = m + 1
        } else {
            r = m - 1
        }
    }

    return l
}
