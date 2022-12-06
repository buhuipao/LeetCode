func rob(nums []int) int {
    // dp[i] = max(dp[i-2]+nums[i], dp[i-1])
    // 压缩空间后，a为dp[i-2], b为dp[i-1]
    // 初始态：a, b 皆为0
    var a, b int
    for i := 0; i < len(nums); i++ {
        b, a = max(a+nums[i], b), b
    }

    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
