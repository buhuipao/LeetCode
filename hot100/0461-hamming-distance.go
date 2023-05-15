func hammingDistance(x int, y int) int {
    v := x ^ y
    var ans int
    for v != 0 {
        v &= (v - 1)
        ans++
    }

    return ans
}
