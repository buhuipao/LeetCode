var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 广度优先遍历
func colorBorder(grid [][]int, row, col, color int) [][]int {
    m, n := len(grid), len(grid[0])
    type point struct{ x, y int }
    borders := []point{}
    originalColor := grid[row][col]
    vis := make([][]bool, m)
    for i := range vis {
        vis[i] = make([]bool, n)
    }

    q := []point{{row, col}}
    vis[row][col] = true
    for len(q) > 0 {
        p := q[0]
        q = q[1:]
        x, y := p.x, p.y
        isBorder := false
        // 上下左右探索
        for _, dir := range dirs {
            nx, ny := x+dir.x, y+dir.y
            if !(0 <= nx && nx < m && 0 <= ny && ny < n && grid[nx][ny] == originalColor) {
                isBorder = true
            } else if !vis[nx][ny] {
                vis[nx][ny] = true
                q = append(q, point{nx, ny})
            }
        }
        if isBorder {
            borders = append(borders, point{x, y})
        }
    }

    for _, p := range borders {
        grid[p.x][p.y] = color
    }
    return grid
}
