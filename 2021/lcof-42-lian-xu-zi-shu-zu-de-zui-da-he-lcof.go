func maxSubArray(nums []int) int {
    var sum int
    // MIN_INT
    ans := ^int(^uint(0) >> 1)
    n := len(nums)
    for i := 0; i < n; i++ {
        // 如果sum大于0继续加
        if sum > 0 {
            sum += nums[i]
        } else {    // 否则应该剔除sum，避免被拖小
            sum = nums[i]
        }
        
        if sum > ans {
            ans = sum
        }
    }
    return ans
}
