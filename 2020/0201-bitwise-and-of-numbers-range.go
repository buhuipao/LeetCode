// 求m，n的公共前缀
func rangeBitwiseAnd(m int, n int) int {
    Max := 1 << 30
    var ans int
    for Max > 0 && (Max & m == Max & n) {
        ans |= Max & m
        Max >>= 1
    }
    return ans
}
