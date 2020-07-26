// DFS
func longestIncreasingPath(matrix [][]int) int {
    // 准备memo
    rows := len(matrix)
    if rows == 0 {
        return 0
    }
    cols := len(matrix[0])
    memo := make([][]int, rows)
    for i := range memo {
        memo[i] = make([]int, cols)
    }
    ans := 1
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if memo[i][j] == 0 {
                memo[i][j] = dfs(matrix, memo, i, j)
                ans = max(ans, memo[i][j])
            }
        }
    }
    return ans
}

func dfs(matrix, memo [][]int, i, j int) int {
    nexts := [][]int{
        []int{0, -1},
        []int{0, 1},
        []int{-1, 0},
        []int{1, 0},
    }
    var v int
    for index := range nexts {
        row, col := i + nexts[index][0], j + nexts[index][1]
        if row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0]) && matrix[i][j] < matrix[row][col] {
            if memo[row][col] == 0 {
                memo[row][col] = dfs(matrix, memo, row, col)
            }
            v = max(v, memo[row][col])
        }
    }
    return 1 + v
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
