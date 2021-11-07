func maxCount(m int, n int, ops [][]int) int {
    a, b := m, n
    for _, op := range ops {
        a, b = min(op[0], a), min(op[1], b)
    }

    return a * b
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
