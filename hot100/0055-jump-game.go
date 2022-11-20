func canJump(nums []int) bool {
    max := 0
    for i := range nums {
        // 如果nums[0~i-1]都不能到达i，直接返回false
        if i > max {
            return false
        }

        // 更新nums[0~i]能达到的最远位置
        if nums[i] + i > max {
            max = nums[i] + i
        }
    }

    return true
}
