func reverse(x int) int {
    if x == 0 {
        return 0
    }
    // 2147483647 -2147483648
    MAX_INT := int(^uint32(0) >> 1)
    MIN_INT := ^int(^uint32(0) >> 1)
    nums := []int{}
    for i := x; i != 0; i = i / 10 {
        nums = append(nums, i % 10)
    }
    ans := 0
    // 先计算前n-1位
    for i := 0; i < len(nums)-1; i++ {
        ans = ans * 10 + nums[i]
    }
    // 最后一位判断是否溢出
    if ans > MAX_INT / 10 || (ans == MAX_INT / 10 && nums[len(nums)-1] > 7) {
        return 0
    }
    if ans < MIN_INT / 10 || (ans == MIN_INT / 10 && nums[len(nums)-1] < -8) {
        return 0
    }
    ans = 10 * ans + nums[len(nums)-1]
    
    return ans
}
