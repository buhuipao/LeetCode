// dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
// 
// 可以优化为一维：
//  pre,        dp[col]
//  dp[col-1],  cur
// dp[col] = min(dp[col-1], dp[col], pre) + 1
func maximalSquare(matrix [][]byte) int {
    row := len(matrix)
    if row == 0 {
        return 0
    }
    col := len(matrix[0])
    dp := make([]int, col+1)

    var ans, pre int
    // 初始化第一行
    for j := 0; j < col; j++ {
        if matrix[0][j] == '1' {
            dp[j+1] = 1
            // 只有一行的特殊情况
            ans = 1
        }
    }

    for i := 1; i < row; i++ {
        pre = 0
        for j := 0; j < col; j++ {
            if matrix[i][j] == '1' {
                dp[j+1], pre = min(min(pre, dp[j]), dp[j+1]) + 1, dp[j+1]
                ans = max(ans, dp[j+1]*dp[j+1])
            } else {
                dp[j+1], pre = 0, 0
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
