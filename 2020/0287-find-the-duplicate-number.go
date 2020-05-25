// 找到链表入环点
func findDuplicate(nums []int) int {
    fast, slow := 0, 0
    for true {
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
