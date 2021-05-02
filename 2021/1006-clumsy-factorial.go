// 此题可以用数学解法，以下为朴素解法：
// (a * b / c + d) - (A * B / C + D)
func clumsy(N int) int {
    var ans int
    if N >= 4 {
        ans += (N * (N-1) / (N-2) + (N-3))
    } else if N == 3 {
        return 6
    } else if N == 2 {
        return 2
    } else {
        return 1
    }
    l := N % 4
    for i := N-4; i > l; i -= 4 {
        ans += ((-1 * i * (i-1) / (i-2)) + i - 3)
    }
    if l == 3 {
        i := 3
        ans += (-1 * i * (i-1) / (i-2))
    } else if l == 2 {
        i := 2
        ans += (-1 * i * (i-1))
    } else if l == 1 {
        i := 1
        ans += (-1 * i)
    }
    return ans
}
