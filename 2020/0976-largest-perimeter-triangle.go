func largestPerimeter(A []int) int {
    n := len(A)
    if n < 3 {
        return 0
    }
    sort.Ints(A)
    for i := n-1; i >= 2; i-- {
        if A[i-2] + A[i-1] > A[i] {
            return A[i-2] + A[i-1] + A[i]
        }
    }
    return 0
}
