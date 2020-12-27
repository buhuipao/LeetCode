func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    m1, m2 := make(map[byte]int), make(map[byte]int)
    for i := range s {
        if _, ok := m1[s[i]]; !ok {
            m1[s[i]] = i
        }
        if _, ok := m2[t[i]]; !ok {
            m2[t[i]] = i
        }
        if m1[s[i]] != m2[t[i]] {
            return false
        }
    }
    return true
}
