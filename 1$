// 模拟法
func spiralOrder(matrix [][]int) []int {
    n := len(matrix)
    if n == 0 {
        return nil
    }
    m := len(matrix[0])
    // 左、右、上、下
    l, r, t, b := 0, m-1, 0, n-1
    ans := make([]int, 0, m * n)
    for true {
        // 往右，上边界+1
        for j := l; j <= r; j++ {
            ans = append(ans, matrix[t][j])
        }
        t++
        if t > b {
            break
        }
        // 往下，右边界-1
        for i := t; i <= b; i++ {
            ans = append(ans, matrix[i][r])
        }
        r--
        if l > r {
            break
        }
        // 往左，下边界-1
        for j := r; j >= l; j-- {
            ans = append(ans, matrix[b][j])
        }
        b--
        if t > b {
            break
        }
        // 往上，左边界+1
        for i := b; i >= t; i-- {
            ans = append(ans, matrix[i][l])
        }
        l++
        if l > r {
            break
        }
    }
    return ans
}
