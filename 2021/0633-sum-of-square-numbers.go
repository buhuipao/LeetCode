func judgeSquareSum(c int) bool {
    if c <= 4 {
        return c != 3
    }
    L, R := 0, int(math.Sqrt(float64(c)))
    for L <= R {
        tmp := L*L + R*R
        if tmp == c {
            return true
        }
        if tmp > c {
            R--
        } else {
            L++
        }
    }
    return false
}
