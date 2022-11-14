// 栈
func isValid(s string) bool {
    m := map[byte]byte{
        '}': '{',
        ']': '[',
        ')': '(',
    }
    var stack []byte
    for i := range s {
        if _, ok := m[s[i]]; !ok {
            stack = append(stack, s[i])
            continue
        }

        if len(stack) == 0 || m[s[i]] != stack[len(stack)-1] {
            return false
        }

        stack = stack[:len(stack)-1]
    }

    return len(stack) == 0
}
