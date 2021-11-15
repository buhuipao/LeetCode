// 等价于：找出1...n中的完全平方数个数，同时为了避免出现精度问题，原数加上0.5
func bulbSwitch(n int) int {
    return int(math.Sqrt(float64(n)+0.5))
}
