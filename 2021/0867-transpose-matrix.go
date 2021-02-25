func transpose(matrix [][]int) [][]int {
    // 初始化结果
    rows, cols := len(matrix), len(matrix[0])
    ans := make([][]int, cols)
    for i := range ans {
        ans[i] = make([]int, rows)
    }
    // 调换i，j
    for i := range matrix {
        for j := range matrix[i] {
            ans[j][i] = matrix[i][j]
        }
    }
    return ans
}
