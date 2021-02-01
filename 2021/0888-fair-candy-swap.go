func fairCandySwap(A []int, B []int) []int {
    var C1, C2 int
    m1, m2 := make(map[int]int), make(map[int]int)
    for i := range A {
        C1 += A[i]
        m1[A[i]] += i
    }
    for i := range B {
        C2 += B[i]
        m2[B[i]] = i 
    }
    C := C1 - C2
    var ans []int
    for k := range m2 {
        tmp := C/2 + k
        if _, ok := m1[tmp]; ok {
            ans = []int{tmp, k}
            break
        }
    }
    return ans
}
