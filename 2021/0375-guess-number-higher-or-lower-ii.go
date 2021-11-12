// dp：f(i,j)= (i≤k≤j)min{k+max(f(i,k−1),f(k+1,j))}
func getMoneyAmount(n int) int {
    f := make([][]int, n+1)
    for i := range f {
        f[i] = make([]int, n+1)
    }
    for i := n - 1; i >= 1; i-- {
        for j := i + 1; j <= n; j++ {
            minCost := math.MaxInt32
            for k := i; k < j; k++ {
                cost := k + max(f[i][k-1], f[k+1][j])
                if cost < minCost {
                    minCost = cost
                }
            }
            f[i][j] = minCost
        }
    }
    return f[1][n]
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}
