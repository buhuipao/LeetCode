func findMinMoves(machines []int) (ans int) {
    tot := 0
    for _, v := range machines {
        tot += v
    }
    n := len(machines)
    if tot%n > 0 {
        return -1
    }
    avg := tot / n
    sum := 0
    for _, num := range machines {
        num -= avg
        sum += num
        ans = max(ans, max(abs(sum), num))
    }
    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}
