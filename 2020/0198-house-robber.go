// dp[i] = max(dp[i-2]+nums[i], dp[i-1])
func rob(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    a, b := 0, nums[0]
    for i := 1; i < n; i++ {
        if a + nums[i] > b {
            a, b = b, a+nums[i]
        } else {
            a, b = b, b
        }
    }
    return b
}
