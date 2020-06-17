func maxScoreSightseeingPair(A []int) int {
    n := len(A)
    if n < 2 {
        return 0
    }
    ans := 0
    preMax := A[0] + 0
    // ans =（A[i] + i + A[j] - j）
    for i := 1; i < n; i++ {
        ans = max(ans, preMax + A[i] - i)
        preMax = max(preMax, A[i] + i)
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
