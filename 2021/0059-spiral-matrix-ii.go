func generateMatrix(n int) [][]int {
    ans := make([][]int, n)
    for i := range ans {
        ans[i] = make([]int, n)
    }
    N := n * n
    a, d := -1, -1
    b, c := n, n
    for x := 1; x <= N; {
        a++
        for j := d + 1; j < b; j++ {  // 从左往右，i继承自a（即i=a+1），j由d、b来限制左右边界
            ans[a][j] = x
            x++
        }
        b--
        for i := a + 1; i < c; i++ {  // 从上往下，j继承自b（即j=b+1），i由a、c限制上下边界
            ans[i][b] = x
            x++
            
        }
        c--
        for j := b - 1; j > d; j-- {  // 从右往左，i继承自c（即i=c-1），j由b、d来限制左右边界
            ans[c][j] = x
            x++
        }
        d++
        for i := c - 1; i > a; i-- {  // 从下往上，j继承自d（即j=d+1），i由c、a限制上下边界
            ans[i][d] = x
            x++
        }
    }
    return ans
}
