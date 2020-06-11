// 单调栈
func dailyTemperatures(T []int) []int {
    ans := make([]int, len(T))
    // 保存索引值
    stack := []int{}
    for i := range T {
        // 弹出温度较小的
        for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
            pre := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            ans[pre] = i - pre
        }
        stack = append(stack, i)
    }
    return ans
}
