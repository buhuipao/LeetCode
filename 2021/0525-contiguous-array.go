// 前缀和的变种，把0变成-1求和即可
func findMaxLength(nums []int) int {
    m := make(map[int]int)
    m[0] = -1
    var ans, sum int
    for i := range nums {
        if nums[i] == 0 {
            sum += -1
        } else {
            sum += 1
        }
        if index, ok := m[sum]; ok {
            ans = max(i-index, ans)
        } else {
            m[sum] = i
        }
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
