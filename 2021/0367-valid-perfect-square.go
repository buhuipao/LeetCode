// 牛顿法
func isPerfectSquare(num int) bool {
    x0 := float64(num)
    for {
        x1 := (x0 + float64(num)/x0) / 2
        if x0-x1 < 1e-6 {
            x := int(x0)
            return x*x == num
        }
        x0 = x1
    }
}
