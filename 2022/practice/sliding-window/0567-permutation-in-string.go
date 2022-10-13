func checkInclusion(s1 string, s2 string) bool {
    needs := make(map[byte]int)
    for i := range s1 {
        needs[s1[i]] += 1
    }

    var l, r, c int
    window := make(map[byte]int)
    for r < len(s2) {
        if _, ok := needs[s2[r]]; ok {
            window[s2[r]] += 1
            if window[s2[r]] == needs[s2[r]] {
                c++
            }
        }

        r++

        // fmt.Println("window:", window, ", c:", c, ", l:", l, ", r:", r)
        for c == len(needs) {
            // fmt.Println("--window:", window, ", c:", c, ", l:", l, ", r:", r)
            if r - l == len(s1) {
                return true
            }

            if window[s2[l]] > 0 {
                window[s2[l]] -= 1
                if window[s2[l]] < needs[s2[l]] {
                    c--
                }
            }

            l++
        }

    }

    return false
}
