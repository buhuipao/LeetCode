func find132pattern(nums []int) bool {
    n := len(nums)
    if n < 3 {
        return false
    }
    // 1、计算num中某个索引位它左边的最小值，其实就是得出某个2值的1值
    mins := make([]int, n)
    mins[0] = nums[0]
    for i := 1; i < n; i++ {
        mins[i] = min(mins[i-1], nums[i])
    }
    // 2、从后往前遍历，探寻某个索引位3的2是否存在
    stack := []int{}
    for i := n-1; i >= 1; i-- {
        if nums[i] > mins[i] {
            for len(stack) > 0 && stack[len(stack)-1] <= mins[i] {
                stack = stack[:len(stack)-1]
            }
            if len(stack) > 0 && stack[len(stack)-1] < nums[i] {
                return true
            }
            stack = append(stack, nums[i])
        }
    } 
    return false
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
