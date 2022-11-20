func maxSubArray(nums []int) int {
    var cur int
    ans := ^int(^uint(0)>>1)
    for i := range nums {
        // 如果cur大于0，对于找出最大sum才有积极意义，否则cur应该被赋值为nums[i]
        if cur > 0 {
            cur += nums[i]
        } else {
            cur = nums[i]
        }

        if cur > ans {
            ans = cur
        }
    }

    return ans
}
