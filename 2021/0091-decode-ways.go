// DP解法
// dp[i] = dp[i-1] (如果当前值不为0，则可继承上一个值) + dp[i-2]（如果前两个值组成的数字小于26，则还可以继承上上个值）
func numDecodings(s string) int {
    n := len(s)
    // a = f[i-2], b = f[i-1], c = f[i]
    a, b, c := 0, 1, 0
    for i := 1; i <= n; i++ {
        c = 0
        if s[i-1] != '0' {
            c += b
        }
        if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
            c += a
        }
        a, b = b, c
    }
    return c
}
