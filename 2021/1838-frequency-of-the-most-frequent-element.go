// 排序 + 双指针
func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)
    ans := 1
    for l, r, total := 0, 1, 0; r < len(nums); r++ {
        total += (nums[r] - nums[r-1]) * (r - l)
        for total > k {
            total -= nums[r] - nums[l]
            l++
        }
        ans = max(ans, r-l+1)
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
