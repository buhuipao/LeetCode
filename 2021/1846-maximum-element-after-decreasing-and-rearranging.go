// 排序
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
    sort.Ints(arr)
    
    arr[0] = 1
    n := len(arr)
    for i := 1; i < n; i++ {
        // arr[i] 只会大于等于 arr[i-1]，于是可以取 min(arr[i], arr[i-1]+1)
        arr[i] = min(arr[i], arr[i-1]+1)
    }

    return arr[n-1]
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
