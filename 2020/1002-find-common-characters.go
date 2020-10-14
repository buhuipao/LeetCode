import "strings"

func commonChars(A []string) []string {
    tmp := helper(A)
    if len(tmp) == 0 {
        return tmp
    }
    return strings.Split(tmp[0], "")
}

func helper(A []string) []string {
    n := len(A)
    if n < 2 {
        return A
    }
    if n > 2 {
        B := helper(A[:n/2])
        C := helper(A[n/2:])
        
        return helper(append(B, C...))
    }
    // 基本情况
    m := make(map[rune]int, 0)
    for _, v := range A[0] {
        m[v] += 1
    }
    ans := make([]rune, 0, len(m))
    for _, v := range A[1] {
        if m[v] > 0 {
            ans = append(ans, v)
            m[v] -= 1
        }
    }
    return []string{string(ans)}
}
