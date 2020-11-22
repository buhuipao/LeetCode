func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    m := make(map[byte]int, 0)
    for i := range s {
        m[s[i]] += 1
        m[t[i]] -= 1
    }
    for k := range m {
        if m[k] != 0 {
            return false
        }
    }
    return true
}
