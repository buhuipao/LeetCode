func maximumGap(nums []int) int {
    n := len(nums)
    if n < 2 {
        return 0
    }
    // 假装桶排序 
    sort.Ints(nums)
    ans := 0
    for i := 1; i < n; i++ {
        tmp := abs(nums[i]-nums[i-1])
        if tmp > ans {
            ans = tmp
        }
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return 0-x
    }
    return x
}
