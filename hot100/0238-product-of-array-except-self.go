func productExceptSelf(nums []int) []int {
    ans := make([]int, len(nums))
    for i := range ans {
        ans[i] = 1
    }

    left := 1
    for i := range nums {
        ans[i] *= left
        left = nums[i] * left
    }

    right := 1
    for i := len(nums)-1; i >= 0; i-- {
        ans[i] *= right
        right *= nums[i]
    }

    return ans
}
