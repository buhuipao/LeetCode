func findMaxAverage(nums []int, k int) float64 {
    sum, n := 0, len(nums)
    for i := 0; i < k; i++ {
        sum += nums[i]
    }
    ans := float64(sum) / float64(k)
    for i := k; i < n; i++ {
        sum = sum - nums[i-k] + nums[i]
        tmp := float64(sum) / float64(k)
        if tmp > ans {
            ans = tmp
        }
    }
    return ans
}
