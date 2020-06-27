func firstMissingPositive(nums []int) int {
    n := len(nums)
    // 利用鸽子洞原理调整数组
    for i := range nums {
        value := nums[i]
        // 不停地找到下一个位置不对的value
        for value <= n && value > 0 && nums[value-1] != value {            
            nums[value-1], value = value, nums[value-1]
        }
    }
    // 查找第一个缺失正数
    for i := range nums {
        if nums[i] != i + 1 {
            return i + 1
        }
    }
    // 或者返回还未出现的第一个数字
    return n + 1
}
