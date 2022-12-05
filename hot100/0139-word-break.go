// dp
// F(x) = bool(F(x-i-j)), 0<j<max(len(word)), word 属于wordDict 
func wordBreak(s string, wordDict []string) bool {
    dict := make(map[string]int, len(wordDict))
    for i := range wordDict {
        dict[wordDict[i]] = len(wordDict[i])
    }

    dp := make([]bool, len(s)+1)
    dp[0] = true
    for i := range s {
        for _, l := range dict {
            if i + 1 < l {
                continue
            }

            // 存在这个单词，且满足条件
            if _, ok := dict[s[i+1-l:i+1]]; ok && dp[i+1-l] {
                dp[i+1] = true
            }
        }
    }

    return dp[len(s)]
}
