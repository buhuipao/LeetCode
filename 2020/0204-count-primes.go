func countPrimes(n int) int {
    if n < 2 {
        return 0
    }
    // 合数数组，默认全是质数
    l := make([]bool, n)
    for i := 2; i * i < n; i++ {
        if !l[i] {
            for j := i*i; j < n; j += i {
                l[j] = true
            }
        }
    }
    var ans int
    for i := range l {
        if !l[i] {
            ans++
        }
    }
    // 减去0、1两个特殊位置
    return ans - 2
}
