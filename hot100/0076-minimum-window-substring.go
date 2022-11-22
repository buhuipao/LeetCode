// 双指针
func minWindow(s string, t string) string {
    count, window := make(map[byte]int), make(map[byte]int)
    for i := range t {
        count[t[i]] += 1
    }

    var l, L, R, valid int
    length := len(s) + 1
    for r := range s {
        if _, ok := count[s[r]]; ok {
            window[s[r]] += 1
            if window[s[r]] == count[s[r]] {
                valid++
            }
        }

        for valid == len(count) {
            if r+1 - l < length {
                L, R = l, r+1
                length = r + 1 - l
            }

            if _, ok := count[s[l]]; ok {
                if window[s[l]] == count[s[l]] {
                    valid--
                }

                window[s[l]] -= 1
            }

            l++
        }

    }

    return s[L:R]
}
