// 快慢指针
func moveZeroes(nums []int)  {
    if len(nums) == 0 {
        return 
    }

    var slow int
    for fast := range nums {
        if nums[fast] != 0 {
            nums[slow] = nums[fast]
            slow++
        }
    }

    for slow < len(nums) {
        nums[slow] = 0
        slow++
    }
}
