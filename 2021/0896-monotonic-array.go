func isMonotonic(A []int) bool {
    n := len(A)
    if n == 0 {
        return true
    }
    flag := (A[n-1] - A[0]) >= 0    // 递增标识
    for i := 1; i < n; i++ {
        if (flag && A[i] < A[i-1]) || (!flag && A[i] > A[i-1]) {
            return false
        }
    }
    return true
}
