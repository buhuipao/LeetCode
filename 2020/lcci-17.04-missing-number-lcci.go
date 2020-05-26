func missingNumber(nums []int) int {
    var ans int
    for i := range nums {
        ans ^= nums[i] ^ i
    }
    ans ^= len(nums)
    return ans
}
