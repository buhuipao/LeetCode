// (rand(x)-1) * rand(y) => [1, x*y]里的随机数
func rand10() int {
    for {
        row := rand7()
        col := rand7()
        idx := (row-1)*7 + col
        if idx <= 40 {
            return 1 + (idx-1)%10
        }
    }
}
