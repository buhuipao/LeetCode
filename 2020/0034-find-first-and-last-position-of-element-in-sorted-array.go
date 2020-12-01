// 利用二分法找左右边界
func searchRange(nums []int, target int) []int {
    ans := []int{-1, -1}
    // 找左边界
    l, r := 0, len(nums)
    for l < r {
        mid := l + (r - l) / 2
        if nums[mid] > target {
            r = mid
        } else if nums[mid] == target {
            // 收紧右边界
            r = mid
            ans[0] = mid
        } else {
            l = mid + 1
        }
    }
    // 找右边界
    l, r = 0, len(nums)
    for l < r {
        mid := l + (r - l) / 2
        if nums[mid] > target {
            r = mid
        } else if nums[mid] == target {
            // 收紧左边界
            ans[1] = mid
            l = mid + 1
        } else {
            l = mid + 1
        }
    }
    return ans
}
