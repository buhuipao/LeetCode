func findMagicIndex(nums []int) int {
    ans := -1
    for i := range nums {
        if nums[i] == i {
            ans = i
            break
        }
    }
    return ans
}
