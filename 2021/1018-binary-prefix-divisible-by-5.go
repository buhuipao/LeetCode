func prefixesDivBy5(A []int) []bool {
    ans := make([]bool, len(A))
    var n int
    for i := range A {
        n = (n << 1 + A[i]) % 5
        ans[i] = n  == 0
    }
    return ans
}
