func hammingDistance(x int, y int) int {
    var ans int
    for i := 0; i < 32; i++ {
        if (x ^ y) & (1 << i) == (1 << i) {
            ans++
        }
    }
    return ans
}
