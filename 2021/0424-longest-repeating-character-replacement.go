func characterReplacement(s string, k int) int {
    cnt := [26]int{}
    maxCnt, left := 0, 0
    for right, ch := range s {
        cnt[ch-'A']++
        maxCnt = max(maxCnt, cnt[ch-'A'])   // 维护left、right间的最大字母计数
        if right-left+1-maxCnt > k {
            cnt[s[left]-'A']--
            left++
        }
    }
    return len(s) - left
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
