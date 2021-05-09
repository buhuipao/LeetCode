func xorOperation(n, start int) (ans int) {
    s, e := start>>1, n&start&1
    ret := sumXor(s-1) ^ sumXor(s+n-1)
    return ret<<1 | e
}

// 相邻四个数字的异或结果为0，且四个为一组进行循环
func sumXor(x int) int {
    switch x % 4 {
    case 0:
        return x
    case 1:
        return 1
    case 2:
        return x + 1
    default:
        return 0
    }
}
