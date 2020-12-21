func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    pre, cur := 0, 0
    for i := 2; i <= n; i++ {
        pre, cur = cur, min(cur+cost[i-1], pre+cost[i-2])
    }
    return cur
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
