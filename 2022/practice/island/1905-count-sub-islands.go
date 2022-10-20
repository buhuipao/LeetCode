var m = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    for i := range grid1 {
        for j := range grid1[i] {
            // 可以理解为地图1中的水为强酸水，盖到地图2上时遇陆地以及其相邻的陆地时，就会腐蚀融化地图2的陆地
            if grid1[i][j] == 0 && grid2[i][j] == 1 {
                dfs(grid2, i, j)
            }
        }
    }

    var ans int
    for i := range grid2 {
        for j := range grid2[i] {
            if grid2[i][j] == 1 {
                ans++
                dfs(grid2, i, j)
            }
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
   for index := range m {
       dfs(grid, i+m[index][0], j+m[index][1])
   }
}
