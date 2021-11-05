// hash 表
func longestSubsequence(arr []int, difference int) (ans int) {
    dp := map[int]int{}
    for _, v := range arr {
        dp[v] = dp[v-difference] + 1
        if dp[v] > ans {
            ans = dp[v]
        }
    }
    return
}
