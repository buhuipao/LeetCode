func minPairSum(nums []int) int {
    sort.Ints(nums)

    n := len(nums)
    var ans int
    for i := 0; i < n/2; i++ {
        ans = max(nums[i]+nums[n-1-i], ans)
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
