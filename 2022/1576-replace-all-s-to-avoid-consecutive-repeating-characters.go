func modifyString(s string) string {
    res := []byte(s)
    n := len(res)
    for i, ch := range res {
        if ch == '?' {
            for b := byte('a'); b <= 'c'; b++ {
                if !(i > 0 && res[i-1] == b || i < n-1 && res[i+1] == b) {
                    res[i] = b
                    break
                }
            }
        }
    }
    return string(res)
}
