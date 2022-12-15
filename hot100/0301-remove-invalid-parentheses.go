// BFS + 剪枝
func removeInvalidParentheses(s string) []string {
    var ans []string
    queue := []string{s}
    added := make(map[string]bool)
    added[s] = true
    for len(queue) != 0 {
        // fmt.Println("queue:", queue)
        str := queue[0]
        queue = queue[1:]
        // 判断是否有答案了，是否还有进一步判断的必要
        // 如果ans中最后一个字符串的长度大于当前的str，证明可以返回了，无需再删字符了
        if len(ans) > 0 && len(ans[len(ans)-1]) > len(str) {
            return ans
        }

        // 判定有效
        if valid(str) {
            ans = append(ans, str)
        }

        // BFS
        for i := range str {
            // 如果不是括号，删除没有意义，所以也不需要加入BFS的队列
            if str[i] != '(' && str[i] != ')' {
                continue
            }

            tmp := str[0:i] + str[i+1:]
            // BFS剪枝
            if added[tmp] {
                continue
            }

            queue = append(queue, tmp)
            added[tmp] = true
        }
    }

    return ans
}

func valid(s string) bool {
    var count int
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
