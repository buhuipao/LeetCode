// DFS
var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}
func updateBoard(board [][]byte, click []int) [][]byte {
    if board[click[0]][click[1]] == 'M' {
        board[click[0]][click[1]] = 'X'
    } else {
        dfs(board, click[0], click[1])
    }
    return board
}

func dfs(board [][]byte, x, y int) {
    cnt := 0
    for i := 0; i < 8; i++ {
        tx, ty := x + dirX[i], y + dirY[i]
        if !inBoard(board, tx, ty) {
            continue
        }
        // 不用判断 M，因为如果有 M 的话游戏已经结束了
        if board[tx][ty] == 'M' {
            cnt++
        }
    }
    if cnt > 0 {
        board[x][y] = byte(cnt + '0')
    } else {
        board[x][y] = 'B'
        for i := 0; i < 8; i++ {
            tx, ty := x + dirX[i], y + dirY[i]
            // 这里不需要在存在 B 的时候继续扩展，因为 B 之前被点击的时候已经被扩展过了
            if !inBoard(board, tx, ty) || board[tx][ty] != 'E' {
                continue
            }
            dfs(board, tx, ty)
        }
    }
}

func inBoard(board [][]byte, x, y int) bool {
    if x < 0 || y < 0 {
        return false
    }
    if x >= len(board) || y >= len(board[0]) {
        return false
    }
    return true
}
