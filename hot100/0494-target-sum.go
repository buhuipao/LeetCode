// DP
func findTargetSumWays(nums []int, target int) int {
    // 求出差值，然后选取nums中应该为负号的数字，使得它们的和为neg
    // neg = (sum - target) / 2 = diff / 2
    var sum int
    for i := range nums {
        sum += nums[i]
    }
    diff := sum - target
    if diff < 0 || diff % 2 != 0 {
        return 0
    }
    
    n, neg := len(nums), diff/2
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, neg+1)
    }
    dp[0][0] = 1
    for i := range nums {
        for j := 0; j <= neg; j++ {
            dp[i+1][j] = dp[i][j]
            if j >= nums[i] {
                dp[i+1][j] += dp[i][j-nums[i]]
            }
        }
    }

    return dp[n][neg]
}
