// 中心扩展法
func countSubstrings(s string) int {
    var ans, left, right int
    n := len(s)
    for i := range s {
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
