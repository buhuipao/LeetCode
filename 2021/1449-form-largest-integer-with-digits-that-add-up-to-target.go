// COPY FROM LEETCODE
func largestNumber(cost []int, target int) string {
    dp := make([][]int, 10)
    from := make([][]int, 10)
    for i := range dp {
        dp[i] = make([]int, target+1)
        for j := range dp[i] {
            dp[i][j] = math.MinInt32
        }
        from[i] = make([]int, target+1)
    }
    dp[0][0] = 0
    for i, c := range cost {
        for j := 0; j <= target; j++ {
            if j < c {
                dp[i+1][j] = dp[i][j]
                from[i+1][j] = j
            } else {
                if dp[i][j] > dp[i+1][j-c]+1 {
                    dp[i+1][j] = dp[i][j]
                    from[i+1][j] = j
                } else {
                    dp[i+1][j] = dp[i+1][j-c] + 1
                    from[i+1][j] = j - c
                }
            }
        }
    }
    if dp[9][target] < 0 {
        return "0"
    }
    ans := make([]byte, 0, dp[9][target])
    i, j := 9, target
    for i > 0 {
        if j == from[i][j] {
            i--
        } else {
            ans = append(ans, '0'+byte(i))
            j = from[i][j]
        }
    }
    return string(ans)
}
