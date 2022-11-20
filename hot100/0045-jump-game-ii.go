func jump(nums []int) int {
    var last, nlast, ans int
    for i := 0; i < len(nums)-1; i++ {
        nlast = max(nlast, nums[i]+i)
        if last == i {
            ans++
            last = nlast
        }
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
