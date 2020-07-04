func longestValidParentheses(s string) int {
    // 栈
    stack := []int{-1}
    ans := 0
    // 两种索引会入栈：等待匹配的左括号索引、充当“分隔符”的右括号索引。
    // 后者也要入栈是因为：当括号匹配光了，栈还需要留一个“垫底”的“参照物”
    for i := range s {
        if s[i] == '(' {
            stack = append(stack, i)
        } else {
            stack = stack[:len(stack)-1]
            // 栈空了则放入，作为之后计算的参照物
            if len(stack) == 0 {
                stack = append(stack, i)
            } else {
                pre := stack[len(stack)-1]
                if i - pre > ans {
                    ans = i - pre
                }
            }
        }
    }

    return ans
}
