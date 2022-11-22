// 一层一层计算矩形面积
func maximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0 {
        return 0
    }

    var ans int
    n := len(matrix[0])
    // 已经在前后设置了哨兵，例如：[0, 1, 0, 1, 0, 0, 0]
    heights := make([]int, n+2)
    for i := range matrix {
        for j := range matrix[i] {
            if matrix[i][j] == '1' {
                heights[j+1] += 1
            } else { // 被中间的0截断，重置高度为0
                heights[j+1] = 0
            }
        }
        ans = max(ans, largestRectangleArea(heights))
    }

    return ans
}

func largestRectangleArea(heights []int) int {
    var ans int
    stack := []int{0}
    for i := 1; i < len(heights); i++ {
        for heights[stack[len(stack)-1]] > heights[i] {
            h := heights[stack[len(stack)-1]]
            w := i - stack[len(stack)-2]-1
            ans = max(h*w, ans)
            stack = stack[:len(stack)-1]
        }

        stack = append(stack, i)
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
