// 单调栈
func dailyTemperatures(temperatures []int) []int {
    ans := make([]int, len(temperatures))
    var stack []int
    for i := range temperatures {
        for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
            pre := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            ans[pre] = i - pre
        }
        stack = append(stack, i)
    }

    return ans
}
