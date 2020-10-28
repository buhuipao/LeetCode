func uniqueOccurrences(arr []int) bool {
    m1 := make(map[int]int)
    for i := range arr {
        m1[arr[i]] += 1
    }
    m2 := make(map[int]bool)
    for _, v := range m1 {
        if m2[v] {
            return false
        }
        m2[v] = true
    }
    return true
}
