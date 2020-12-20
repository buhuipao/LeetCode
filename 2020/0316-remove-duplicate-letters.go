func removeDuplicateLetters(s string) string {
    // 先统计s中字母频率，确保不会有子母被错误的跳过和落下
    left := [26]int{}
    for _, ch := range s {
        left[ch-'a']++
    }
    stack := []byte{}
    inStack := [26]bool{}
    for i := range s {
        ch := s[i]
        // 如果当前字母不在栈里则需要需要处理，否则直接忽略
        if !inStack[ch-'a'] {
            // 为确保字典序尽可能小，必须要尝试弹出当前的大字符
            for len(stack) > 0 && ch < stack[len(stack)-1] {
                // 如过准备弹出的字符是此字符的最后一个，则必须保留，跳出循环
                last := stack[len(stack)-1] - 'a'
                if left[last] == 0 {
                    break
                }
                stack = stack[:len(stack)-1]
                inStack[last] = false
            }
            stack = append(stack, ch)
            inStack[ch-'a'] = true
        }
        // 减去计数
        left[ch-'a']--
    }
    return string(stack)
}
