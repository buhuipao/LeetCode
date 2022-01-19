func containsNearbyDuplicate(nums []int, k int) bool {
    m := make(map[int]int, len(nums)/2)
    for i := range nums {
        if v, ok := m[nums[i]]; ok {
            if i - v <= k {
                return true
            }
        }

        m[nums[i]] = i
    }

    return false
}
