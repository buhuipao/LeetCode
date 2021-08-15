// 下面是leetcode给出的dp做法，其实仔细思考这其实是一个数学问题
const mod int = 1e9 + 7
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func findPaths(m, n, maxMove, startRow, startColumn int) (ans int) {
    dp := make([][][]int, maxMove+1)
    for i := range dp {
        dp[i] = make([][]int, m)
        for j := range dp[i] {
            dp[i][j] = make([]int, n)
        }
    }
    dp[0][startRow][startColumn] = 1
    for i := 0; i < maxMove; i++ {
        for j := 0; j < m; j++ {
            for k := 0; k < n; k++ {
                count := dp[i][j][k]
                if count > 0 {
                    for _, dir := range dirs {
                        j1, k1 := j+dir.x, k+dir.y
                        if j1 >= 0 && j1 < m && k1 >= 0 && k1 < n {
                            dp[i+1][j1][k1] = (dp[i+1][j1][k1] + count) % mod
                        } else {
                            ans = (ans + count) % mod
                        }
                    }
                }
            }
        }
    }
    return
}
