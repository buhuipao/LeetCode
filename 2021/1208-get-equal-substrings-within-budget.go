func equalSubstring(s string, t string, maxCost int) (maxLen int) {
    n := len(s)
    accDiff := make([]int, n+1)
    for i, ch := range s {
        accDiff[i+1] = accDiff[i] + abs(int(ch)-int(t[i]))
    }
    for i := 1; i <= n; i++ {
        start := sort.SearchInts(accDiff[:i], accDiff[i]-maxCost)
        maxLen = max(maxLen, i-start)
    }
    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
