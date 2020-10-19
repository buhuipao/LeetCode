// 使用栈
func backspaceCompare(S string, T string) bool {
    helper := func(s string) string {
        stack := make([]byte, 0, len(s)/4)
        for i := range s {
            if s[i] != '#' {
                stack = append(stack, s[i])
            } else {
                if len(stack) > 0 {
                    stack = stack[:len(stack)-1]
                }
            }
        }
        return string(stack)
    }
    
    return helper(S) == helper(T)
}
