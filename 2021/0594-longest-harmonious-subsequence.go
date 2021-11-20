func findLHS(nums []int) int {
    m := make(map[int]int, len(nums)/2)
    var ans int
    for _, n := range nums {
        m[n]++
        if m[n-1] != 0 {
            ans = max(m[n]+m[n-1], ans)
        }

        if m[n+1] != 0 {
            ans = max(m[n]+m[n+1], ans)
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
