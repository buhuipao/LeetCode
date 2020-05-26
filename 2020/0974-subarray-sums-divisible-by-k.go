// 累加和
func subarraysDivByK(A []int, K int) int {
    var ans, sum int
    m := make(map[int]int, 0)
    m[0] = 1
    for i := range A {
        sum += A[i]
        A[i] = sum
        tmp := sum%K
        if tmp < 0 {
            tmp += K
        }
        ans += m[tmp]
        m[tmp] += 1
        
    }
    return ans
}
