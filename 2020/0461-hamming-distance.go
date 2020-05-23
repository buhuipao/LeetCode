// 异或之后统计1的个数
func hammingDistance(x int, y int) int {
    z := x ^ y
    var ans int
    for z != 0 {
        ans++
        z &= z - 1
    }
    return ans
}
