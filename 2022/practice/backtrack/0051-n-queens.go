func solveNQueens(n int) [][]string {
    var ans [][]string

    backtrack(n, 0, &ans, nil)

    return ans
}

func backtrack(n, cur int, paths *[][]string, path []string) {
    if cur == n {
        *paths = append(*paths, path)
        return
    }

    bs := genS(n)
    for j := 0; j < n; j++ {
        // 剪枝
        if !isValid(n, cur, j, &path) {
            continue
        }

        // 做出选择
        tmp := make([]string, len(path))
        copy(tmp, path)
        bs[j] = 'Q'
        tmp = append(tmp, string(bs))
        // 进入下一层决策
        backtrack(n, cur+1, paths, tmp) 

        // 撤销选择
        bs[j] = '.'
    }
}


func genS(n int) []byte {
    ret := make([]byte, n)
    for i := 0; i < n; i++ {
        ret[i] = '.'
    } 
    return ret
}

func isValid(n, row, col int, path *[]string) bool {
    if len(*path) == 0 {
        return true
    }

    // 检查正上方
    for i := range *path {
        if (*path)[i][col] == 'Q' {
            return false
        }
    }

    // 检查左上方
    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if i >= 0 && j >= 0 && (*path)[i][j] == 'Q' {
            return false
        }
    }

    // 检查右上方
    for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
        if i >= 0 && j < n && (*path)[i][j] == 'Q' {
            return false
        }
    }

    return true
}
