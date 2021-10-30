func singleNumber(nums []int) []int {
    // 找到两个不同值的异或值
    x := 0
    for i := range nums {
        x ^= nums[i]
    }
    if x == 0 {
        return nil
    }
    
    // 找到两个值的第一个低差异位
    y := x - x & (x - 1)
    var a, b int
    for i := range nums {
        if nums[i] & y != 0 {
            a ^= nums[i]
        } else {
            b ^= nums[i]
        }
    }

    return []int{a, b}
}
