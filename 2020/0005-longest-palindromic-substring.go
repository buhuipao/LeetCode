func longestPalindrome(s string) string {
    n := len(s)
    if n == 0 {
        return ""
    }
    L, R := 0, 0
    for i := range s {
        left1, right1 := findPalindrome(s, i, i)
        left2, right2 := findPalindrome(s, i, i+1)
        if right2 - left2 > right1 - left1 {
            left1, right1 = left2, right2
        }
        if right1 - left1 > R - L {
            L, R = left1, right1
        }
    }
    return s[L:R+1]
}


func findPalindrome(s string, left, right int) (int, int) {
    n := len(s)
    for left >= 0 && right < n && s[left] == s[right] {
            left--
            right++
    }
    return left+1, right-1
}
