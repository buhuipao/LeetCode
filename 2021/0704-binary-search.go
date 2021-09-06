func search(nums []int, target int) int {
    for l, r := 0, len(nums)-1; l <= r; {
        m := l + (r -l) / 2
        if nums[m] == target {
            return m
        }

        if nums[m] < target {
            l = m + 1
        } else {
            r = m - 1
        }
    }
    return -1
}
