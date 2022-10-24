func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    
    ans1, ans2 := robRange(nums, 0, n-2), robRange(nums, 1, n-1)
    if ans1 > ans2 {
        return ans1
    }
    
    return ans2
}

func robRange(nums []int, l, r int) int {
    var dp, dp1, dp2 int
    for i := l; i <= r; i++ {
        if dp2 + nums[i] > dp1 {
            dp = dp2 + nums[i]
        }

        dp1, dp2 = dp, dp1
    } 

    return dp
}
