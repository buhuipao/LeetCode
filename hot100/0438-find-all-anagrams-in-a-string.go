func findAnagrams(s string, p string) []int {
    needs := make(map[byte]int)
    for i := range p {
        needs[p[i]] +=1
    }

    var l, r, c int
    var ans []int
    window := make(map[byte]int, len(needs))
    for r < len(s) {
        if _, ok := needs[s[r]]; ok {
            window[s[r]] += 1
            if window[s[r]] == needs[s[r]] {
                c++
            }
        }
        r++

        for c == len(needs) {
            if r - l == len(p) { // 符合条件
                ans = append(ans, l)
            }

            if window[s[l]] > 0 {
                window[s[l]] -= 1
                if window[s[l]] < needs[s[l]] {
                    c--
                }
            }

            l++
        }

    }

    return ans
}
