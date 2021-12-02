var medals = map[int]string{
    0: "Gold Medal",
    1: "Silver Medal",
    2: "Bronze Medal",
}

func findRelativeRanks(score []int) []string {
    n := len(score)
    m := make(map[int]int, n)
    for i, s := range score {
        m[s] = i
    }

    sort.Ints(score)
    ans := make([]string, n)
    var v string
    for i, s := range score {
        if n - i <= 3 {
            v = medals[n-1-i]
        } else {
            v = fmt.Sprintf("%d", n-i)
        }

        ans[m[s]] = v
    }

    return ans
}
