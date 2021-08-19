func reverseVowels(s string) string {
    t := []byte(s)
    n := len(t)
    i, j := 0, n-1
    for i < j {
        for i < n && !strings.Contains("aeiouAEIOU", string(t[i])) {
            i++
        }
        for j > 0 && !strings.Contains("aeiouAEIOU", string(t[j])) {
            j--
        }
        if i < j {
            t[i], t[j] = t[j], t[i]
            i++
            j--
        }
    }
    return string(t)
}
