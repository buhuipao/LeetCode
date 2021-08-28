func runningSum(nums []int) []int {
    var tmp int
    for i := range nums {
        tmp += nums[i]
        nums[i] = tmp
    }
    
    return nums
}
