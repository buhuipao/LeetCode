// 中心扩散法
func longestPalindrome(s string) string {
    var L, R int
    var l1, r1, l2, r2 int
    for i := range s {
        // 考虑奇偶情况，进行两次查找
        l1, r1 = findPalindrom(s, i, i)
        l2, r2 = findPalindrom(s, i, i+1)

        if r1 - l1 > r2 - l2 {
            r2, l2 = r1, l1
        }

        if r2 - l2 > R - L {
            R, L = r2, l2
        }
    }

    return s[L:R+1]
}

func findPalindrom(s string, l, r int) (int, int) {
    for l >= 0 && r < len(s) && s[l] == s[r] {
        l--
        r++
    }

    return l+1, r-1
}
