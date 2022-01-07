func maxDepth(s string) (ans int) {
    size := 0
    for _, ch := range s {
        if ch == '(' {
            size++
            if size > ans {
                ans = size
            }
        } else if ch == ')' {
            size--
        }
    }
    return
}
