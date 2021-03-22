func hammingWeight(num uint32) int {
    var ans int
    for num != 0 {
        ans++
        num &= (num-1)
    }
    return ans
}
