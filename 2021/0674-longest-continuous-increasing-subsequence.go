func findLengthOfLCIS(nums []int) int {
    if len(nums) < 1 {
        return 0
    }
    n, tmp, ans := len(nums), 1, 1
    for i := 1; i < n; i++ {
        if nums[i] > nums[i-1] {
            tmp++
        } else {
            tmp = 1
        }
        if ans < tmp {
            ans = tmp
        }
    }
    return ans
}
