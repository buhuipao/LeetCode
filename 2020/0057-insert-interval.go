func insert(intervals [][]int, newInterval []int) [][]int {
    i, n := 0, len(intervals)    
    ans := make([][]int, 0, n+1)
    // 第一步：把与newInterval无交集的数组推入ans
    for i < n && intervals[i][1] < newInterval[0] {
        ans = append(ans, intervals[i])
        i++
    }
    // 第二步：合并newInterval
    for i < n && intervals[i][0] <= newInterval[1] {
        newInterval[0] = min(intervals[i][0], newInterval[0])
        newInterval[1] = max(intervals[i][1], newInterval[1])
        i++
    }
    ans = append(ans, newInterval)
    // 第三步；追加剩余数组
    for i < n {
        ans = append(ans, intervals[i])
        i++
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
