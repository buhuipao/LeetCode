func removeDuplicates(nums []int) int {
    L := 0
    for i := range nums {
        if i > 0 && nums[i-1] == nums[i] {
            continue
        }
        nums[L] = nums[i]
        L++
    }
    return L
}
