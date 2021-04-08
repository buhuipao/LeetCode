func findMin(nums []int) int {
    n := len(nums)
    if n == 0 {
        return -1
    }
    L, R := 0, n-1
    for L < R {
        M := L + (R-L)/2
        if nums[M] > nums[R] {
            L = M + 1
        } else {
            R = M
        }
    }
    return nums[L]
}
