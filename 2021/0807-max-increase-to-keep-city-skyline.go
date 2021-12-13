// 将每一个值改成这个值所在行最大值、列最大值里更小的那一个
func maxIncreaseKeepingSkyline(grid [][]int) (ans int) {
    n := len(grid)
    rowMax := make([]int, n)
    colMax := make([]int, n)
    for i, row := range grid {
        for j, h := range row {
            rowMax[i] = max(rowMax[i], h)
            colMax[j] = max(colMax[j], h)
        }
    }
    for i, row := range grid {
        for j, h := range row {
            ans += min(rowMax[i], colMax[j]) - h
        }
    }
    return
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
