func maxPower(s string) int {
    if s == "" {
        return 0
    }

    ans, tmp := 1, 1
    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1] {
            tmp++
        } else {
            tmp = 1
        }

        ans = max(ans, tmp)
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
