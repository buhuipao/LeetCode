func removeDuplicates(nums []int) int {
    var L int
    count := 0
    for i := range nums {
        if i > 0 && nums[i] == nums[i-1] {
            count++
            if count > 2 {
                continue
            }
        } else {
            count = 1
        }
        nums[L] = nums[i]
        L++
    }
    return L
}
