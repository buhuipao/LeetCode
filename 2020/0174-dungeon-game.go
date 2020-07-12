// 原地DP
func calculateMinimumHP(dungeon [][]int) int {
    m := len(dungeon)
    if m == 0 {
        return 0
    }
    n := len(dungeon[0])
    // 原地DP定义：到达终点时当前点需要保留的最低能量
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            if i == m-1 && j == n-1 {
                // 如果终点大于0，只需要保留1的能量到达终点即可，不需要多余能量
                if dungeon[i][j] > 0 {
                    dungeon[i][j] = 1
                // 如果终点小于0，那么就需要使得最终到达终点时，能量等于1
                } else {
                    dungeon[i][j] = 1 + (0 - dungeon[i][j])
                }
            // 最下边一行只能从左往右走，同时要保证当前状态大于等于1
            } else if i == m-1 {
                dungeon[i][j] = max(dungeon[i][j+1]-dungeon[i][j], 1)
            // 最右边一列只能从上往下走，同时要保证当前状态大于等于1
            } else if j == n-1 {
                dungeon[i][j] = max(dungeon[i+1][j]-dungeon[i][j], 1)
            // 选择往下或者往右走
            } else {
                dungeon[i][j] = max(1, min(dungeon[i][j+1], dungeon[i+1][j])-dungeon[i][j])
            }
        }
    }
    return dungeon[0][0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
