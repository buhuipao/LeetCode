// 排序+stack
func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return nil
    }

    // 排序
    sort.Sort(Intervals(intervals))
    var ans [][]int
    var n int
    // 合并
    for i := range intervals {
        ans = append(ans, intervals[i])
        n = len(ans)
        // 判断ans[n-1]与ans[n-2]是否有交集
        for n >= 2 && ans[n-1][0] <= ans[n-2][1] {
            ans[n-2][0], ans[n-2][1] = min(ans[n-1][0], ans[n-2][0]), max(ans[n-1][1], ans[n-2][1])
            ans, n = ans[:n-1], n-1
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

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

type Intervals [][]int

func (I Intervals) Len() int {
    return len(I)
}

func (I Intervals) Less(i, j int) bool {
    return I[i][0] < I[j][0]
}

func (I Intervals) Swap(i, j int) {
    I[i], I[j] = I[j], I[i]
}
