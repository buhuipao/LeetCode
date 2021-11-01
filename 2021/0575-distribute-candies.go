func distributeCandies(candyType []int) int {
    n := len(candyType)
    m := make(map[int]int)
    for _, c := range candyType {
        m[c] += 1
    }

    return min(n/2, len(m))
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
