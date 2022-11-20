func search(nums []int, target int) int {
    n := len(nums)
    if n == 0 {
        return -1
    }

    l, r := 0, n-1
    for l <= r {
        mid := l + (r - l) / 2
        if nums[mid] == target {
            return mid
        }

        // mid左边有序
        if nums[0] <= nums[mid] {
            if nums[0] <= target && target < nums[mid] {
                r = mid - 1
            } else {
                l = mid + 1
            }
        } else {
            if nums[mid] < target && target <= nums[n-1] {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
    }

    return -1
}
