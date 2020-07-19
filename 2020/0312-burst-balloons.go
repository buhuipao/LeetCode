// dp, copy from https://leetcode-cn.com/problems/burst-balloons/solution/dong-tai-gui-hua-tao-lu-jie-jue-chuo-qi-qiu-wen-ti/
func maxCoins(nums []int) int {
    n := len(nums)
    nums = append(nums, 1)
    nums = append([]int{1}, nums...)
    // dp[i][j] = dp[i][k] + dp[k][j] + nums[i]*nums[k]*nums[j]
    dp := make([][]int, n+2)
    for i := range dp {
        dp[i] = make([]int, n+2)
    }
    // i 应该从下往上
    for i := n; i >= 0; i-- {
        // j 应该从左往右
        for j := 0; j <= n+1; j++ {
            // 最后戳破的气球是哪个？
            for k := i+1; k < j; k++ {
                dp[i][j] = max(dp[i][j], dp[i][k] + dp[k][j] + nums[i]*nums[k]*nums[j])
            }
        }
    }

    return dp[0][n+1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
