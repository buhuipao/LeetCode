func pivotIndex(nums []int) int {
    n := len(nums)
    if n == 0 {
        return -1
    }
    sums := make([]int, n)
    sums[n-1] = nums[n-1]
    for i := n-2; i >= 0; i-- {
        sums[i] += nums[i] + sums[i+1]
    }
    var total int
    for i := 0; i < n; i++ {
        total += nums[i]
        if total == sums[i] {
            return i
        }
    }
    return -1
}
