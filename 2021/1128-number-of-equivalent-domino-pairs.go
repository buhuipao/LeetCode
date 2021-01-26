func numEquivDominoPairs(dominoes [][]int) int {
    m := make(map[int]int)
    var ans int
    for i := range dominoes {
        a, b := dominoes[i][0], dominoes[i][1]
        if a > b {
            a, b = b, a
        }
        s := a * 10 + b
        ans += m[s]
        m[s] += 1
    }
    return ans
}
