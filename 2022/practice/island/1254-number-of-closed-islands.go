func closedIsland(grid [][]int) int {
    if len(grid) == 0 {
        return 0
    }

   rows, cols := len(grid), len(grid[0])
    // 淹没左右
    for i := 0; i < rows; i++ {
        dfs(grid, i, 0)
        dfs(grid, i, cols-1)
    }
    // 淹没上下
    for i := 0; i < cols; i++ {
        dfs(grid, 0, i)
        dfs(grid, rows-1, i) 
    }

    var ans int
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 0 {
                ans++
            }
            dfs(grid, i, j)
        }
    }
    
    return ans
}

func dfs(grid [][]int, i, j int) {
   rows, cols := len(grid), len(grid[0])
   if i < 0 || j < 0 || i >= rows || j >= cols {
       return
   }

   if grid[i][j] == 1 {
       return
   }

   grid[i][j] = 1
   for _, v := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
       dfs(grid, i+v[0], j+v[1])
   }
}
