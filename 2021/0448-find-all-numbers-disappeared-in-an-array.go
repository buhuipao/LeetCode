func findDisappearedNumbers(nums []int) []int {
    n := len(nums)
    for _, v := range nums {
        v = (v - 1) % n // 数组内已有值的索引
        nums[v] += n    // 标记目标索引值的值
    }
    ans := make([]int, 0)
    for i, v := range nums {
        if v <= n {
            ans = append(ans, i+1)  // 找出未被标记的索引
        }
    }
    return ans
}
