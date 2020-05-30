// 单调栈
func largestRectangleArea(heights []int) int {
    // 添加左边的哨兵，它永远不会弹出，避免栈为空
    // 添加右边的哨兵，保障最小栈完全弹出
    heights = append([]int{0}, heights...)
    heights = append(heights, 0)

    var ans int
    n := len(heights)
    stack := []int{0}
    for i := 1; i < n; i++ {
        for heights[i] < heights[stack[len(stack)-1]] {
            curH := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            // 计算宽度时必须计算至前一位
            curW := i - stack[len(stack)-1] - 1
            tmp := curH * curW
            if tmp > ans {
                ans = tmp
            }
        }
        stack = append(stack, i)
    }
    return ans
}
