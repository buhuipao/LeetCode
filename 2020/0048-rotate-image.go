func rotate(matrix [][]int)  {
    n := len(matrix)
    // 水平翻转
    for i := 0; i < n/2; i++ {
        for j := 0; j < n; j++ {
            matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
        }
    }
    // 对角翻转
    for i := 0; i < n; i++ {
        for j := 0; j < i; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
}
