func countBinarySubstrings(s string) int {
    // 0011100110 => [2, 3, 2, 2, 1]
    var pre byte
    var preCount, curCount, ans int
    for i := range s {
        if s[i] != pre {
            ans += min(preCount, curCount)
            preCount, curCount = curCount, 0
        }
        curCount++
        pre = s[i]
    }
    // 加上最后一次计算
    ans += min(preCount, curCount)

    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
