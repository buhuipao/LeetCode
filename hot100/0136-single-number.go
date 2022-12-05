func singleNumber(nums []int) int {
    var ans int
    for i := range nums {
        ans ^= nums[i]
    }

    return ans
}
