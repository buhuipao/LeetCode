// 栈
func isValid(s string) bool {
    stack := []rune{}
    m := map[rune]rune {
        '}': '{',
        '{': '0',
        ']': '[',
        '[': '0',
        ')': '(',
        '(': '0',

    }
    for _, v := range s {
        if m[v] == '0' {
            stack = append(stack, v)
            continue
        }
        if len(stack) == 0 || m[v] != stack[len(stack)-1] {
            return false
        } 
        stack = stack[:len(stack)-1]
    }
    return len(stack) == 0
}
