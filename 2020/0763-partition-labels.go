func partitionLabels(S string) []int {
    // 统计
    count := make(map[rune]int, 26)
    for _, s := range S {
        count[s] += 1
    }
    pre, sum := -1, 0
    ans := make([]int, 0)
    for i, s := range S {
        // 增加路径，并重置
        sum += count[s]
        count[s] = 0
        
        // 当步长为1时证明该走的路都已经走过路，符合要求了
        if sum == 1 {
            ans = append(ans, i-pre)
            pre = i
        }
        sum--
    }
    return ans
}
