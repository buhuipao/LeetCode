func lengthOfLongestSubstring(s string) int {
    var l, r, ans int
    window := make(map[byte]int)
    for r < len(s) {
        c := s[r]
        window[c] +=1
        r++

        for window[c] > 1 {
            window[s[l]] -= 1
            l++
        }

        if r - l > ans {
            ans = r - l
        }
    }

    return ans
}
