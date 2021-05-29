func subarraySum(nums []int, k int) (ans int) {
    mp := map[int]int{0: 1}
    for i, pre := 0, 0; i < len(nums); i++ {
        pre += nums[i]
        if _, ok := mp[pre-k]; ok {
            ans += mp[pre-k]
        }
        mp[pre]++
    }
    return
}

func numSubmatrixSumTarget(matrix [][]int, target int) (ans int) {
    for i := range matrix { // 枚举上边界
        sum := make([]int, len(matrix[0]))
        for _, row := range matrix[i:] { // 枚举下边界
            for c, v := range row {
                sum[c] += v // 更新每列的元素和
            }
            ans += subarraySum(sum, target)
        }
    }
    return
}
