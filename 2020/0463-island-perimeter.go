func islandPerimeter(grid [][]int) int {
    if len(grid) == 0 {
        return 0
    }
    var ans int
    n, m := len(grid), len(grid[0])
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == 1 {
                ans += length(grid, i, j)
            }
        }
    }
    return ans
}

// 计算单个岛屿的周长，旁边多一个岛屿周长减1
func length(grid [][]int, i, j int) int {
    n, m := len(grid), len(grid[0])
    l := 4
    if i-1 >= 0 && grid[i-1][j] == 1 {
        l--
    }
    if i+1 < n && grid[i+1][j] == 1 {
        l--
    }
    if j-1 >= 0 && grid[i][j-1] == 1 {
        l--
    }
    if j+1 < m && grid[i][j+1] == 1 {
        l--
    }
    return l
}
