// 差分数组
func carPooling(trips [][]int, capacity int) bool {
    ans := make([]int, 1001)
    for i := range trips {
        ans[trips[i][1]] += trips[i][0]
        if trips[i][2] < 1001 {
            ans[trips[i][2]] -= trips[i][0]
        }
    }

    for i := 1; i < 1001; i++ {
        if ans[i-1] > capacity {
            return false
        }

        ans[i] += ans[i-1]
    }

    return true
}
