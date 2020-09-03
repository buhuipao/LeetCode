func solveNQueens(n int) [][]string {
    ans := make([][]string, 0)
    path := make([][]byte, n)
    for i := 0; i < n; i++ {
        path[i] = make([]byte, n)
        for j := 0; j < n; j++ {
            path[i][j] = '.'
        }
    }
    backtrack(path, 0, n, &ans)

    return ans
}

func backtrack(path [][]byte, i, n int, ans *[][]string) {
    // 终止条件
    if i == n {
        tmp := make([]string, n)
        for i := 0; i < n; i++ {
            tmp[i] = string(path[i])
        }
        *ans = append(*ans, tmp)
        return
    }
    for j := 0; j < n; j++ {
        // 剪枝
        if !validate(path, i, j) {
            continue
        }
        // 做出选择
        path[i][j] = 'Q'
        backtrack(path, i+1, n, ans)
        // 撤销选择
        path[i][j] = '.'
    }
}

// 判断当前位置如果放下一个皇后，是否合法
func validate(path [][]byte, row, col int) bool {
    n := len(path)

    // 判断竖直方向
    for i := row-1; i >= 0 ; i-- {
        if path[i][col] == 'Q' {
            return false
        }
    }

    // 判断正斜线和反斜线方向
    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if path[i][j] == 'Q' {
            return false
        }
    }
    for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
        if path[i][j] == 'Q' {
            return false
        }
    }
    return true
}


func codyPath(p1 [][]byte) [][]byte {
    n := len(p1)
    p2 := make([][]byte, n)
    for i := 0; i < n; i++ {
        p2[i] = make([]byte, n)
        copy(p2[i], p1[i])
    }
    return p2
}
