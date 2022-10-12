func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    var slow int
    for fast := range nums {
        if nums[fast] != nums[slow] {
            // 右移慢指针
            slow++
            nums[slow]= nums[fast]
        }
    }

    return slow+1
}
