func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }

    var ans int
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == '1' {
                ans++
            }

            // 淹没这块岛屿的所有陆地
            dfs(grid, i, j)
        }
    }
    
    return ans
}

func dfs(grid [][]byte, i, j int) {
    if len(grid) == 0 {
        return 
    }

    row, col := len(grid), len(grid[0])
    // 检查边界
    for i < 0 || j < 0 || i >= row || j >= col {
        return
    }

    // 是否已经淹没
    if grid[i][j] == '0' {
        return
    }

    grid[i][j] = '0'
    for _, v := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
        dfs(grid, i+v[0], j+v[1])
    }
}
