func lengthOfLIS(nums []int) int {
    n := len(nums)
    if n < 2 {
        return n
    }

    var dp []int
    dp = append(dp, nums[0])
    for i := 1; i < n ; i++ {
        if nums[i] > dp[len(dp)-1] {
            dp = append(dp, nums[i])
            continue
        }

        // 找到合适的位置替换为nums[index]，然后可以维持一个递增的数组，以便有更大的数字时可以直接追加到最后
        l, r := 0, len(dp)-1
        for l <= r {
            m := l + (r - l) / 2
            if dp[m] < nums[i] {
                l = m + 1
            } else if dp[m] < nums[i] {
                r = m - 1
            } else {
                r = m - 1
            }
        }
        dp[l] = nums[i]
    }

    return len(dp)
}
