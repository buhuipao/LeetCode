func maxAreaOfIsland(grid [][]int) int {
    var ans, tmp int
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 0 {
                continue
            }

            tmp = dfs(grid, i, j)
            if tmp > ans {
                ans = tmp
            }
        }
    }

    return ans
}

func dfs(grid [][]int, i, j int) int {
    rows, cols := len(grid), len(grid[0])
    if i < 0 || j < 0 || i >= rows || j >= cols {
        return 0 
    }

    if grid[i][j] == 0 {
        return 0 
    }

    grid[i][j] = 0
    var cnt int
    for _, v := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
        cnt += dfs(grid, i+v[0], j+v[1])
    }

    return cnt+1
}
