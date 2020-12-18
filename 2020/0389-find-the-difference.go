func findTheDifference(s string, t string) byte {
    m := make(map[byte]int)
    for i := range s {
        m[s[i]] -= 1
        m[t[i]] += 1
    }
    m[t[len(t)-1]] += 1
    for i := range m {
        if m[i] == 1 {
            return i
        }
    }
    return '0'
}
