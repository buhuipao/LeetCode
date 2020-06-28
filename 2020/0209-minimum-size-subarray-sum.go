// 滑动窗口
func minSubArrayLen(s int, nums []int) int {
    var ans, sum int
    L, R := 0, -1
    n := len(nums)
    ans = n + 1
    for i := range nums {
        // 向右扩张
        sum += nums[i]
        R++
        // 收紧左边
        for L <= R && sum >= s {
            if ans > R - L + 1 {
                ans = R - L + 1
            }
            sum -= nums[L]
            L++
        }
    }
    if ans == n + 1 {
        return 0
    }
    return ans
}
