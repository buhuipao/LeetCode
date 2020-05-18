// dp[i] = max(dp[i-1]*nums[i], nums[i])
func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    ans := ^int(^uint(0)>>1)
    myMax, myMin := 1, 1
    for i := range nums {
        // 如果nums[i]小于0，调换最大和最小值
        if nums[i] < 0 {
            myMax, myMin = myMin, myMax
        }
        if myMax * nums[i] > nums[i] {
            myMax *= nums[i]
        } else {
            myMax = nums[i]
        }
        if myMin * nums[i] < nums[i] {
            myMin *= nums[i]
        } else {
            myMin = nums[i]
        }
        if myMax > ans {
            ans = myMax
        }
    }
    return ans
}
