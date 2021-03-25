// 只需要记录当前的行和列
func setZeroes(matrix [][]int)  {
    r := len(matrix)
    if r == 0 {
        return
    }
    c := len(matrix[0])

    // 标记第一行和第一列是否有0
    var rowFlag, colFlag bool
    for j := 0; j < c; j++ {
        if matrix[0][j] == 0 {
            rowFlag = true
            break
        }
    }
    for i := 0; i < r; i++ {
        if matrix[i][0] == 0 {
            colFlag = true
            break
        }
    }
    // 
    for i := 1; i < r; i++ {
        for j := 1; j < c; j++ {
            // 把标记放到第一行和第一列
            if matrix[i][j] == 0 {
                matrix[0][j] = 0
                matrix[i][0] = 0
            }
        }
    }
    // 真正的标记，注意忽略第一行和第一列
    for j := 1; j < c; j++ {
        if matrix[0][j] == 0 {
            for i := 1; i < r; i++ {
                matrix[i][j] = 0
            }
        }
    }
    for i := 1; i < r; i++ {
        if matrix[i][0] == 0 {
            for j := 1; j < c; j++ {
                matrix[i][j] = 0
            }
        }
    }
    // 处理第一行和第一列
    if rowFlag {
        for j := 0; j < c; j++ {
            matrix[0][j] = 0
        }
    }
    if colFlag {
        for i := 0; i < r; i++ {
            matrix[i][0] = 0
        }
    }
}
