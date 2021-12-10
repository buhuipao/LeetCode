func shortestCompletingWord(licensePlate string, words []string) string {
    count := make(map[byte]int)
    s := strings.ToLower(licensePlate)
    for i := range s {
        if s[i] >= 'a' && s[i] <= 'z' {
            count[s[i]]++
        }
    }

    min := 16
    var ans string
    for _, w := range words {
        m := make(map[byte]int)
        for i := range w {
            m[w[i]]++
        }

        flag := true
        for k, v := range count {
            if m[k] < v {
                flag = false
                break
            }
        }

        if flag && len(w) < min {
            min = len(w)
            ans = w
        }
    }

    return ans
}
