func wordPattern(pattern string, s string) bool {
    strs := strings.Split(s, " ")
    n := len(pattern)
    if n != len(strs) {
        return false
    }
    m1, m2 := make(map[byte]int), make(map[string]int)
    for i := 0; i < n; i++ {
        if _, ok := m1[pattern[i]]; !ok {
            m1[pattern[i]] = i
        }
        if _, ok := m2[strs[i]]; !ok {
            m2[strs[i]] = i
        }
        if m1[pattern[i]] != m2[strs[i]] {
            return false
        }
    }
    return true
}
