func flipAndInvertImage(A [][]int) [][]int {
    if len(A) == 0 {
        return nil
    }
    cols := len(A[0])
    for i := range A {
        for j := 0; j < cols/2; j++ {
            A[i][j], A[i][cols-1-j] = A[i][cols-1-j]^1, A[i][j]^1
        }
        // 对于奇数列的矩阵，需要特殊处理中间那一列
        if cols % 2 == 1 {
            A[i][cols/2] ^= 1
        }
    }
    return A
}
