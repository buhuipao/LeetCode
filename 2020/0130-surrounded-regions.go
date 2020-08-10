// DFS
func solve(board [][]byte)  {
    rows := len(board)
    if rows == 0 {
        return
    }
    cols := len(board[0])
    // 找到所有和边缘相接的'O'，并替换成'B'
    for i := 0; i < rows; i++ {
        if board[i][0] == 'O' {
            dfs(board, i, 0)
        }
        if board[i][cols-1] == 'O' {
            dfs(board, i, cols-1)
        }
    }
    for j := 0; j < cols; j++ {
        if board[0][j] == 'O' {
            dfs(board, 0, j)
        }
        if board[rows-1][j] == 'O' {
            dfs(board, rows-1, j)
        }
    }
    // 替换
    for i := 0; i < rows; i++{
        for j := 0; j < cols; j++ {
            if board[i][j] == 'B' {
                board[i][j] = 'O'
            } else if board[i][j] == 'O' {
                board[i][j] = 'X'
            }
        }
    }
}

func dfs(board [][]byte, row, col int) {
    if board[row][col] == 'O' {
        board[row][col] = 'B'
    }
    nexts := [][]int{
        []int{0, 1},
        []int{0, -1},
        []int{1, 0},
        []int{-1, 0},
    }
    for i := range nexts {
        newR, newC := row + nexts[i][0], col + nexts[i][1]
        if newR >= 0 && newR < len(board) && newC >= 0 && newC < len(board[0]) && board[newR][newC] == 'O' {
            dfs(board, newR, newC)
        }
    }
}
