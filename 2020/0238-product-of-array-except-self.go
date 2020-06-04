func productExceptSelf(nums []int) []int {
    left := 1
    n := len(nums)
    ans := make([]int, n)
    for i := 0; i < n; i++ {
        ans[i], left = left, left * nums[i]
    }
    right := 1
    for i := n-1; i >= 0; i-- {
        ans[i] *= right
        right *= nums[i]
    }
    return ans
}
