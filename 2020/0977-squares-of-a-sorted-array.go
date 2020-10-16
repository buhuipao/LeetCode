func sortedSquares(a []int) []int {
    n := len(a)
    ans := make([]int, n)
    i, j := 0, n-1
    for pos := n - 1; pos >= 0; pos-- {
        // 取出较大者
        if v, w := a[i]*a[i], a[j]*a[j]; v > w {
            ans[pos] = v
            i++
        } else {
            ans[pos] = w
            j--
        }
    }
    return ans
}
