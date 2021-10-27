// 逐步试探 + 队列
func removeInvalidParentheses(s string) []string {
    ans := make([]string, 0)
    m1, m2 := make(map[string]bool, 0), make(map[string]bool, 0)
    queue := []string{s}
    preL := len(s)
    for len(queue) > 0 {
        str := queue[0]
        queue = queue[1:len(queue)]
        // 没有加入 && 合法
        if _, ok := m1[str]; !ok && isValid(str) {
            if len(ans) == 0 {
                ans = append(ans, str)
                m1[str] = true
                preL = len(str)
            } else {
                if preL == len(str) {
                    ans = append(ans, str)
                    m1[str] = true
                    preL = len(str)
                } else {
                    return ans
                }
            }
        }
        for i := range str {
            if str[i] != '(' && str[i] != ')' {
                continue
            }
            // 避免往队列里加入重复字符串
            newS := str[:i] + str[i+1:]
            if _, ok := m2[newS]; !ok {
                queue = append(queue, newS)
                m2[newS] = true
            }
        }
    }
    return ans
}

func isValid(s string) bool {
    count := 0
    for i := range s {
        if s[i] == '(' {
            count++
        } else if s[i] == ')' {
            count--
        }
        if count < 0 {
            return false
        }
    }
    return count == 0
}
