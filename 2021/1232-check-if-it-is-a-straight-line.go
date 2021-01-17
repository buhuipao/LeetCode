func checkStraightLine(coordinates [][]int) bool {
    n := len(coordinates)
    if n < 2 {
        return true
    }
    x1 := float64(coordinates[1][0] - coordinates[0][0])
    y1 := float64(coordinates[1][1] - coordinates[0][1])
    // 算出斜率
    // 为了避免处理除数为0的情况，这里可以把除法转为乘法: b/a == y/x => ay == xb
    if x1 == 0 {    // 同一竖轴上
        for i := 2; i < n; i++ {
            a, b := coordinates[i-1], coordinates[i]
            if a[0] - b[0] != 0 {
                return false
            }
        } 
    } else if y1 == 0 { // 同一横轴上
        for i := 2; i < n; i++ {
            a, b := coordinates[i-1], coordinates[i]
            if a[1] - b[1] != 0 {
                return false
            }
        }
    } else {
        v := y1 / x1
        for i := 2; i < n; i++ {
            a, b := coordinates[i-1], coordinates[i]
            if float64(b[1]-a[1]) / float64(b[0]-a[0]) != v {
                return false
            }
        }
    }
    return true
}
