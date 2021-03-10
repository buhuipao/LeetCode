// 使得剩下的n-k个连续节点的和最小即可
func maxScore(cardPoints []int, k int) int {
    n := len(cardPoints)
    K := n - k
    var sum, tmp int
    for i := range cardPoints {
        sum += cardPoints[i]
        if i < K {
            tmp += cardPoints[i]
        }
    }
    ans := sum - tmp
    for i := K; i < n; i++ {
        tmp = tmp - cardPoints[i-K] + cardPoints[i]
        ans = max(sum - tmp, ans)
    }
    return ans
}

func max(a, b int)int {
    if a > b {
        return a
    }
    return b
}
