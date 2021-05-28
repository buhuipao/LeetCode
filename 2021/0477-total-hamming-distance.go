func totalHammingDistance(nums []int) (ans int) {
    n := len(nums)
    for i := 0; i < 30; i++ {
        c := 0
        for _, val := range nums {
            c += val >> i & 1
        }
        ans += c * (n - c)
    }
    return
}
