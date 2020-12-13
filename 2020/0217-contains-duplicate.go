func containsDuplicate(nums []int) bool {
    m := make(map[int]bool)
    for _, n := range nums {
        if m[n] {
            return true
        }
        m[n] = true
    }
    return false
}
