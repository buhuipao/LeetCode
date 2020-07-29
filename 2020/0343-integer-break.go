import "math"

func integerBreak(n int) int {
    if n <= 3 {
        return n - 1
    }
    a, b := float64(n / 3), float64(n % 3)
    if b == 1 {
        return int(math.Pow(3, a-1)) * 4
    }
    if b == 2 {
        return int(math.Pow(3, a)) * 2
    }
    return int(math.Pow(3, a))
}
