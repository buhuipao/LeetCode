// 双指针
func moveZeroes(nums []int)  {
    l, r, n := 0, 0, len(nums)
    // 确保nums[l, r] 全都为0，nums[0:l]都为非零数
    for r < n {
        if nums[r] != 0 {
            nums[l], nums[r] = nums[r], nums[l]
            l++
        }
        r++
    }
}
