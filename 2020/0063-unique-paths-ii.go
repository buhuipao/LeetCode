// DP: dp[i][j] = dp[i-1][j] + dp[i][j-1] when ob[i][j] != 1
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    rows := len(obstacleGrid)
    if rows == 0 {
        return 0
    }
    cols := len(obstacleGrid[0])
    // 初始化
    dp := make([][]int, rows)
    for i := range dp {
        dp[i] = make([]int, cols)
    }
    // 初始化最左和最上的一列
    for i := 0; i < rows; i++ {
        if obstacleGrid[i][0] == 1 {
            break
        }
        dp[i][0] = 1
    }
    for j := 0; j < cols; j++ {
        if obstacleGrid[0][j] == 1 {
            break
        }
        dp[0][j] = 1
    }
    // 计算
    for i := 1; i < rows; i++ {
        for j := 1; j < cols; j++ {
            if obstacleGrid[i][j] != 1 {
                dp[i][j] = dp[i-1][j] + dp[i][j-1]
            }
        }
    }
    return dp[rows-1][cols-1]
}
