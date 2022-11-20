// 原地DP
func minPathSum(grid [][]int) int {
    m := len(grid)
    if m == 0 {
        return 0
    }

    n := len(grid[0])
    // 计算左边
    for i := 1; i < m; i++ {
        grid[i][0] += grid[i-1][0]
    }

    // 计算上边
    for j := 1; j < n; j++ {
        grid[0][j] += grid[0][j-1]
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            // 取左边和上边的较小值
            grid[i][j] += min(grid[i-1][j], grid[i][j-1])
        }
    }

    return grid[m-1][n-1]
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
