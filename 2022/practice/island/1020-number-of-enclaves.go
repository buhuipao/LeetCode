func numEnclaves(grid [][]int) int {
    if len(grid) == 0 {
        return 0
    }

   rows, cols := len(grid), len(grid[0])
   for i := 0; i < rows; i++ {
       dfs(grid, i, 0)
       dfs(grid, i, cols-1)
   }
   for j := 0; j < cols; j++ {
       dfs(grid, 0, j)
       dfs(grid, rows-1, j)
   }

   var ans int
   for i := range grid {
       for j := range grid[i] {
            ans += grid[i][j]
       }
   }

   return ans
}

func dfs(grid [][]int, i, j int) {
   rows, cols := len(grid), len(grid[0])
   if i < 0 || j < 0 || i >= rows || j >= cols {
       return
   }

   if grid[i][j] == 0 {
       return
   }

   grid[i][j] = 0
   for _, v := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
       dfs(grid, i+v[0], j+v[1])
   }
}
