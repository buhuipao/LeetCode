func numSubarraysWithSum(nums []int, goal int) int {
    // 初始化
    m := make(map[int]int, len(nums)/8)
    
    // 统计
    var sum, ans int
    for i := range nums {
        m[sum] += 1
        sum += nums[i]
        ans += m[sum-goal]
    }    

    return ans
}
