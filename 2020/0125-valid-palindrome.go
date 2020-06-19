func isPalindrome(s string) bool {
    l, r := 0, len(s) - 1
    for l < r {
        for true {
            if l >= len(s) {
                return true
            }
            if !helper(s[l]) {
                l++
            } else {
                break
            } 
        }
        for true {
            if r < 0 {
                return true
            }
            if !helper(s[r]) {
                r--
            } else {
                break
            }
        }
        a, b := s[l], s[r]
        if a > 'Z' {
            a -= ('a' - 'A')
        }
        if b > 'Z' {
            b -= ('a' - 'A')
        }
        if a != b {
            return false
        }
        l++
        r--
    }
    return true
}

func helper(a byte) bool {
    return (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9')
}
