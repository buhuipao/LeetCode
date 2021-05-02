// dp
// 第一个和最后一个不能同时取到，于是最终答案是：max(rob(nums[:n-1]), rob(nums[1:n])
func rob(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return nums[0]
    }
    return max(helper(nums[0:n-1]), helper(nums[1:n]))
}

// dp[x] = max(dp[x-2]+nums[x], dp[x-1])
func helper(nums []int) int {
    n := len(nums)
    dp := []int{0, 0}
    for i := 0; i < n; i++ {
        dp[0], dp[1] = dp[1], max(dp[0]+nums[i], dp[1])
    } 
    return dp[1]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
