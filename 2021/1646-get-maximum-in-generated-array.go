// 找规律
func getMaximumGenerated(n int) (ans int) {
    if n == 0 {
        return
    }
    nums := make([]int, n+1)
    nums[1] = 1
    for i := 2; i <= n; i++ {
        nums[i] = nums[i/2] + i%2*nums[i/2+1]
    }
    for _, v := range nums {
        ans = max(ans, v)
    }
    return
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
