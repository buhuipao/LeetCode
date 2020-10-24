func videoStitching(clips [][]int, t int) int {
    const inf = math.MaxInt64 - 1
    dp := make([]int, t+1)
    for i := range dp {
        dp[i] = inf
    }
    dp[0] = 0
    for i := 1; i <= t; i++ {
        for _, c := range clips {
            l, r := c[0], c[1]
            // 1）若能剪出子区间 [l,i]，则可以从 dp[l] 转移到 dp[i]
            // 2）并且符合最小数目要求
            if l < i && i <= r && dp[l]+1 < dp[i] {
                dp[i] = dp[l] + 1
            }
        }
    }
    if dp[t] == inf {
        return -1
    }
    return dp[t]
}
