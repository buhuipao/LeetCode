// dp
// F(x) = bool(F(x-i-j)), 0<j<max(len(word)), word 属于wordDict 
func wordBreak(s string, wordDict []string) bool {
    dict := make(map[string]int, len(wordDict))
    for _, w := range wordDict {
        dict[w] = len(w)
    }
    n := len(s)
    dp := make([]bool, n+1)
    dp[0] = true
    for i := 0; i < n; i++ {
        // 尝试不同的字典长度
        for _, l := range dict {
            if i + 1 < l {
                continue
            }
            if _, ok := dict[s[i+1-l:i+1]]; ok && dp[i+1-l] {
                    dp[i+1] = true
            }
        }
    }
    return dp[n]
}
