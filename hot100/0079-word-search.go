// 回溯
func exist(board [][]byte, word string) bool {
    if len(word) == 0 {
        return true
    }
    m := len(board)
    if m == 0 {
        return false
    }
    n := len(board[0])
    path := make([][]int, m)
    // 全部初始化为0
    for i := 0; i < m; i++ {
        path[i] = make([]int, n)
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            // 找到world的第一个单词
            if board[i][j] == word[0] {
                // 进入回溯
                if backtrack(path, board, i, j, word, 0) {
                    return true
                }
            }
        }
    }
    return false
}
func backtrack(path [][]int, options [][]byte, i, j int, target string, index int) bool {
    // 第一个字符就不同无需探索
    if options[i][j] != target[index] {
        return false
    }
    // 所找字符串的最后一个字符满足，结束回溯
    if len(target) - 1 == index {
        return true
    }
    // 锁住i/j位置
    path[i][j] = 1
    // 向右
    if j + 1 < len(options[0]) && path[i][j+1] != 1 && backtrack(path, options, i, j+1, target, index+1) {
        return true
    } 
    // 向下
    if i + 1 < len(options) && path[i+1][j] != 1 && backtrack(path, options, i+1, j, target, index+1) {
        return true
    }
    // 向左
    if j - 1 >= 0 && path[i][j-1] != 1 && backtrack(path, options, i, j-1, target, index+1) {
        return true
    }
    // 向上
    if i - 1 >= 0 && path[i-1][j] != 1 && backtrack(path, options, i-1, j, target, index+1) {
        return true
    }
    // 撤销锁定
    path[i][j] = 0
    return false
}
