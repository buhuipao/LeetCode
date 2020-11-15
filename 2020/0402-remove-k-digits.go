func removeKdigits(num string, k int) string {
    stack := []byte{}
    for i := range num {
        d := num[i]
        // 123463 => 1233，把“46”弹出
        for k > 0 && len(stack) > 0 && d < stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
            k--
        }
        stack = append(stack, d)
    }
    // 排除单调上升栈剩余的几位
    // 同时剪掉前导零
    ans := strings.TrimLeft(string(stack[:len(stack)-k]), "0")
    if ans == "" {
        return "0"
    }
    return ans

}
