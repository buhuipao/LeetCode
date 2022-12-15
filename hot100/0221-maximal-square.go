func maximalSquare(matrix [][]byte) int {
    row := len(matrix)
    if row == 0 {
        return 0
    }

    col := len(matrix[0])
    dp := make([]int, col+1)

    var ans, pre int
    // 初始化第一行
    for j := range matrix[0] {
        if matrix[0][j] == '1' {
            dp[j+1] = 1
            ans = 1
        }
    }

    for i := 1; i < row; i++ {
        pre = 0
        for j := 0; j < col; j++ {
            if matrix[i][j] == '1' {
                pre, dp[j+1] = dp[j+1], min(min(dp[j], dp[j+1]), pre)+1
                ans = max(ans, dp[j+1]*dp[j+1])
            } else {
                pre, dp[j+1] = dp[j+1], 0
            }
        }
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
