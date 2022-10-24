func rob(nums []int) int {
    var dp1, dp2 int
    // 默认为0
    dp, n := 0, len(nums)
    // 从最后一家开始抢
    for i := n-1; i >= 0; i-- {
        // 如果抢了nums[i]，则需要加上nums[i-2:]之后的收益
        // 然后跟不抢nums[i-1]的收益作对比
        if dp2 + nums[i] > dp1 {
            dp = dp2 + nums[i]
        } else {
            dp = dp1
        }
        dp1, dp2 = dp, dp1
    }

    return dp
}
