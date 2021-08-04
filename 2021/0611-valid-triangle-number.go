// 排序 + 双指针 （两边之和大于第三边）
func triangleNumber(nums []int) (ans int) {
    n := len(nums)
    sort.Ints(nums)
    
    for i, v := range nums {
        k := i
        for j := i + 1; j < n; j++ {
            for k+1 < n && nums[k+1] < v+nums[j] {
                k++
            }
            ans += max(k-j, 0)
        }
    }
    return
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
