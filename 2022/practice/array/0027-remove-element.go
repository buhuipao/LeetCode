func removeElement(nums []int, val int) int {
    if len(nums) == 0 {
        return 0
    }

    var slow int
    for fast := range nums {
        if nums[fast] != val {
            nums[slow] = nums[fast]
            slow++
        }
    }

    return slow
}
