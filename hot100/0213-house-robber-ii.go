// 因为第一间和最后一间相邻了，所以有两种抢法：1）不抢最后一间；2）不抢第一间
func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    
    v1 := robX(nums, 0, len(nums)-2)
    v2 := robX(nums, 1, len(nums)-1)
    return max(v1, v2)
}

func robX(nums []int, l, r int) int {
    var a, b int
    for i := l; i <= r; i++ {
        a, b = b, max(nums[i]+a, b)
    }

    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
