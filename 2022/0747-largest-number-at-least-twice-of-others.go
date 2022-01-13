func dominantIndex(nums []int) int {
    var max, pre, ans int
    for i := range nums {
        if nums[i] > max {
            max, pre = nums[i], max
            ans = i
            continue
        }

        if nums[i] > pre {
            pre = nums[i]
        }
    }

    if max / 2 >= pre {
        return ans
    }

    return -1
}
