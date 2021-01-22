func addToArrayForm(A []int, K int) []int {
    if K == 0 {
        return A
    }
    B := []int{}
    for K != 0 {
        B = append(B, K % 10)
        K /= 10
    }
    n := len(B)
    for i := 0; i < n/2; i++ {
        B[i], B[n-1-i] = B[n-1-i], B[i]
    }
    m := len(A)
    ans, c, v := []int{}, 0, 0
    for i, j := n-1, m-1; c != 0 || i >= 0 || j >= 0; i, j = i-1, j-1 {
        var v1, v2 int
        if i >= 0 {
            v1 = B[i]
        }
        if j >= 0 {
            v2 = A[j]
        }
        v, c = (v1+v2+c)%10, (v1+v2+c)/10
        ans = append(ans, v)
    }
    l := len(ans)
    for i := 0; i < l/2; i++ {
        ans[i], ans[l-1-i] = ans[l-1-i], ans[i]
    }
    return ans
}
