func twoSum(nums []int, target int) []int {
    m := make(map[int]int, 0)
    for i, v := range nums {
        if j, ok := m[target-v]; ok {
            return []int{j, i}
        }
        m[v] = i
    }

    return nil
}
