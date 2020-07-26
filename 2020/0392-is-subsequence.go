// 双指针
func isSubsequence(s string, t string) bool {
    m, n := len(t), len(s)
    var i, j int
    for i < m && j < n {
        for i < m && s[j] != t[i] {
            i++
        }
        if i < m && j < n {
            j++
            i++
        }
    }
    return j == n
}
