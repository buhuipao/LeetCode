// COPY FROM LEETCODE
func minCost(houses []int, cost [][]int, m, n, target int) int {
    const inf = math.MaxInt64 / 2 // 防止整数相加溢出

    // 将颜色调整为从 0 开始编号，没有被涂色标记为 -1
    for i := range houses {
        houses[i]--
    }

    // dp 所有元素初始化为极大值
    dp := make([][][]int, m)
    for i := range dp {
        dp[i] = make([][]int, n)
        for j := range dp[i] {
            dp[i][j] = make([]int, target)
            for k := range dp[i][j] {
                dp[i][j][k] = inf
            }
        }
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if houses[i] != -1 && houses[i] != j {
                continue
            }

            for k := 0; k < target; k++ {
                for j0 := 0; j0 < n; j0++ {
                    if j == j0 {
                        if i == 0 {
                            if k == 0 {
                                dp[i][j][k] = 0
                            }
                        } else {
                            dp[i][j][k] = min(dp[i][j][k], dp[i-1][j][k])
                        }
                    } else if i > 0 && k > 0 {
                        dp[i][j][k] = min(dp[i][j][k], dp[i-1][j0][k-1])
                    }
                }

                if dp[i][j][k] != inf && houses[i] == -1 {
                    dp[i][j][k] += cost[i][j]
                }
            }
        }
    }

    ans := inf
    for _, res := range dp[m-1] {
        ans = min(ans, res[target-1])
    }
    if ans == inf {
        return -1
    }
    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
