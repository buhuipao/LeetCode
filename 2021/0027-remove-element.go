func removeElement(nums []int, val int) int {
    var L int
    for i := range nums {
        if nums[i] == val {
            continue
        }
        nums[L], nums[i] = nums[i], nums[L]
        L++
    }
    return L
}
