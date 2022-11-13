func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i := range nums {
        if j, ok := m[target-nums[i]]; ok {
            return []int{j, i}
        }

        m[nums[i]] = i
    }

    return nil
}
