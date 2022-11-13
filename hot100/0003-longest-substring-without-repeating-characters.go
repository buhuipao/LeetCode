// 滑动窗口
func lengthOfLongestSubstring(s string) int {
    var l, r, ans int
    var c byte
    window := make(map[byte]int)
    for r < len(s) {
        c = s[r]
        window[c] += 1
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
