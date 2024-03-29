// 按照题意模拟
func reverseStr(s string, k int) string {
    t := []byte(s)
    for i, n := 0, len(s); i < n; i += 2 * k {
        sub := t[i:min(i+k, len(s))]
        for j, n := 0, len(sub); j < n/2; j++ {
            sub[j], sub[n-1-j] = sub[n-1-j], sub[j]
        }
    }
    
    return string(t)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
