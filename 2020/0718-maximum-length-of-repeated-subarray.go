// DP
func findLength(A []int, B []int) int {
    n, m := len(A), len(B)
    if m == 0 || n == 0 {
        return 0
    }
    // 初始化一维的dp
    dp := make([]int, m+1)
    var ans int
    // 从右上角往下开始计算
    for i := 1; i <= n; i++ {
        for j := m; j >= 1; j-- {
            if A[i-1] == B[j-1] {
                dp[j] = dp[j-1] + 1
            } else {
                dp[j] = 0
            }
            if dp[j] > ans {
                ans = dp[j]
            }
        }
    }
    return ans
}
