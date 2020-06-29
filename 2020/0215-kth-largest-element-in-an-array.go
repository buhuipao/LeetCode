func findKthLargest(nums []int, k int) int {
    m := len(nums) - k
    l, r := 0, len(nums)-1
    for true {
        i := helper(nums, l, r)
        if i == m {
            break
        } else if i < m {
            l = i+1
        } else {
            r = i-1
        }
    }
    return nums[m]
}

func helper(nums []int, l, r int) int {
    v := nums[l]
    var i, j int
    // 循环过程中：nums[l]等于v，nums[l+1:j+1]小于v，nums[j+1:r+1]大于等于v
    for i, j = l+1, l; i <= r; i++ {
        if nums[i] < v {
            j++
            nums[i], nums[j] = nums[j], nums[i]
        }
    }
    // 交换后：nums[l:j]小于v，nums[j:r+1]大于等于v
    nums[l], nums[j] = nums[j], nums[l]
    return j
}
