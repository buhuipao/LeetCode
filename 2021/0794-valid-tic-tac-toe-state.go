func win(board []string, p byte) bool {
    for i := 0; i < 3; i++ {
        if board[i][0] == p && board[i][1] == p && board[i][2] == p ||
            board[0][i] == p && board[1][i] == p && board[2][i] == p {
            return true
        }
    }
    return board[0][0] == p && board[1][1] == p && board[2][2] == p ||
        board[0][2] == p && board[1][1] == p && board[2][0] == p
}

func validTicTacToe(board []string) bool {
    oCount, xCount := 0, 0
    for _, row := range board {
        oCount += strings.Count(row, "O")
        xCount += strings.Count(row, "X")
    }
    return !(oCount != xCount && oCount != xCount-1 ||
        oCount != xCount && win(board, 'O') ||
        oCount != xCount-1 && win(board, 'X'))
}
