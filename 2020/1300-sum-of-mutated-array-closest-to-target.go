func findBestValue(arr []int, target int) int {
    sort.Ints(arr)
    n := len(arr)
    if n == 0 {
        return 0
    }
    var ans, sum int
    // 默认取数组最后一个值
    ans = arr[n-1]
    for i := range arr {
        c := (target - sum) / (n - i)
        // 如果当前数组值大于预期平均值，证明所求值在arr[i]之下
        if c < arr[i] {
            // 找到最准确的值
            f := float64(target - sum) / float64(n - i)
            if f - float64(c) > 0.5 {
                return c+1
            } else {
                return c
            }
        }
        sum += arr[i]
    }
    return ans
}
