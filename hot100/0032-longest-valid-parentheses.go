func longestValidParentheses(s string) int {
    // 放入索引-1作为哨兵
    stack := []int{-1}
    var ans int
    for i := range s {
        if s[i] == '(' {
            stack = append(stack, i)
            continue
        }

        // 弹出栈顶的'('，使其暴露这个左括号左边相邻的右括号索引值
        // 1) 如果栈为空，证明之前一直是')'，没有过'('，那本次操作可以看作是用当前'('替换之前旧的
        // 2) 如果栈不为空，弹出的就是一个'('，我们直接基于当前栈顶的索引值，计算中间的跨度后，与ans相比较即可
        stack = stack[:len(stack)-1]
        if len(stack) == 0 {
            stack = append(stack, i)
        } else {
            pre := stack[len(stack)-1]
            if i - pre > ans {
                ans = i - pre
            }
        }
    }

    return ans
}
