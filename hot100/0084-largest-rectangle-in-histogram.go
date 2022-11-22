// 单调栈
func largestRectangleArea(heights []int) int {
    if len(heights) == 0 {
        return 0
    }

    // 前后添加哨兵
    // 添加左边的哨兵，它永远不会弹出，避免栈为空
    // 添加右边的哨兵，保障最小栈完全弹出
    heights = append([]int{0}, heights...)
    heights = append(heights, 0)

    var ans int
    stack := []int{0}
    for i := 1; i < len(heights); i++ {
        // 确保栈是单调递增，所以需要先将大于heights[i]的都弹出
        for heights[stack[len(stack)-1]] > heights[i] {
            h := heights[stack[len(stack)-1]]
            // 宽度需要计算至stack内前一根柱子
            w := i - stack[len(stack)-2] - 1
            if h * w > ans {
                ans = h * w
            }

            stack = stack[:len(stack)-1]
        }


        stack = append(stack, i)
    }

    return ans
}
