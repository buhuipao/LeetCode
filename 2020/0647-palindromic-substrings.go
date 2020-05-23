// 朴素做法：中心扩展法
func countSubstrings(s string) int {
    n := len(s)
    if n == 0 {
        return 0
    }
    var ans int
    var left, right int
    for i := 0; i < n; i++ {
        left, right = i, i
        for left >= 0 && right < n && s[left] == s[right] {
            ans++
            left--
            right++
        }
        left, right = i, i+1
        for left >= 0 && right < n && s[left] == s[right] {
            ans++
            left--
            right++
        }
    }

    return ans
}

// 高效做法：马拉车

