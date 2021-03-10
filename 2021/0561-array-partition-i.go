func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    var ans int
    for i := range nums {
        if i % 2 == 0 {
            ans += nums[i]
        }
    }
    return ans
}
