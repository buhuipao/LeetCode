func canPartition(nums []int) bool {
    // 检查
    n := len(nums)
    if n < 2 {
        return false
    }

    sum, max := 0, 0
    for _, v := range nums {
        sum += v
        if v > max {
            max = v
        }
    }
    if sum%2 != 0 {
        return false
    }

    target := sum / 2
    if max > target {
        return false
    }

    // dp
    dp := make([]bool, target+1)
    // 初始化
    dp[0], dp[nums[0]] = true, true
    for i := 1; i < n; i++ {
        v := nums[i]
        // 为了保证dp[j] = dp[j] || dp[j-v] 能用到上一排的dp[j-v]，需要从大到小更新dp值
        for j := target; j >= v; j-- {
            dp[j] = dp[j] || dp[j-v]
        }
    }

    return dp[target]
}
