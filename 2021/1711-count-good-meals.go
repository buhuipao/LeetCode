// Hash
func countPairs(deliciousness []int) (ans int) {
    maxVal := deliciousness[0]
    for _, val := range deliciousness[1:] {
        if val > maxVal {
            maxVal = val
        }
    }
    // 最大sum，只需要试探小于等于此值的就行
    maxSum := maxVal * 2

    cnt := map[int]int{}
    for _, val := range deliciousness {
        for sum := 1; sum <= maxSum; sum <<= 1 {
            ans += cnt[sum-val] // 只要取两个数的值，只要不遗漏就行，可以这么统计
        }
        cnt[val]++
    }

    return ans % (1e9 + 7)
}
