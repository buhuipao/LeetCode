// 快慢指针
func findDuplicate(nums[]int) int {
    var fast, slow int
    for {
        fast, slow = nums[nums[fast]], nums[slow]
        if fast == slow {
            break
        }
    }

    fast = 0
    for fast != slow {
        fast, slow = nums[fast], nums[slow]
    }

    return fast
}
