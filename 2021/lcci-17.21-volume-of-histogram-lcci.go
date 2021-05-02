func trap(height []int) int {
    stack := []int{}
    n := len(height)
    var ans int
    for i := 0; i < n; i++ {
        for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
            m := len(stack)
            if m > 1 {
                w := i - stack[m-2] - 1
                h := min(height[i], height[stack[m-2]]) - height[stack[m-1]]
                ans += w * h
            }
            stack = stack[:m-1]
        }
        stack = append(stack, i)        
    }
    return ans
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
