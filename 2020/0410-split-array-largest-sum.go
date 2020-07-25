func splitArray(nums []int, m int) int {
    var L, R int
    for i := range nums {
        L = max(L, nums[i])
        R += nums[i]
    }
    for L < R {
        // 以mid为分界和去切分数组
        mid := L + (R - L) / 2
        tmp, count := 0, 1
        for i := range nums {
            if tmp + nums[i] > mid {
                count++
                tmp = nums[i]
            } else{
                tmp += nums[i]
            }
        }
        // 如果按照mid去切分，切出来的份数比m多，则证明mid太小了不符合要求
        if count > m {
            L = mid + 1
        } else {
            R = mid
        }
    }
    return L
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
