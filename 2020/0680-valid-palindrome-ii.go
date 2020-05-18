// 页的二次机会置换算法
func validPalindrome(s string) bool {
    l, r := 0, len(s)-1
    return helper(&s, l, r, 1)
}

func helper(s *string, l, r, count int) bool {
    if l >= r {
        return true
    }
    for l < r {
        if (*s)[l] != (*s)[r] {
            if count < 1 {
                return false
            } else {
                // 尝试删去左边或者右边的字符
                return helper(s, l+1, r, 0) || helper(s, l, r-1, 0)
            }
        } else {
            l++
            r--
        }
    }
    return true
}
