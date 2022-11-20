func searchRange(nums []int, target int) []int {
    ans := []int{-1, -1}
    // 寻找左边界
    l, r := 0, len(nums)-1
    for l <= r {
        mid := l + (r - l) / 2
        if nums[mid] == target {
            ans[0] = mid 
            r = mid - 1
        } else if nums[mid] < target {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }

    // 寻找右边界
    l, r = 0, len(nums)-1
    for l <= r {
        mid := l + (r - l) / 2
        if nums[mid] == target {
            ans[1] = mid
            l = mid + 1
        } else if nums[mid] < target {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }

    return ans
}
