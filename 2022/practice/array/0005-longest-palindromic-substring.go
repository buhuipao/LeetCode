// 中心扩散法
func longestPalindrome(s string) string {
    var L, R int
    for i := range s {
        l1, r1 := findPalindrome(s, i, i)
        l2, r2 := findPalindrome(s, i, i+1)

        if r1 - l1 < r2 - l2 {
            l1, r1 = l2, r2
        }

        if R - L < r1 - l1 {
            L, R  = l1, r1 
        }
    }

    return s[L:R+1]
}

func findPalindrome(s string, l, r int) (int, int) {
    for l >= 0 && r < len(s) && s[l] == s[r] {
        l--
        r++
    }

    return l+1, r-1
}
