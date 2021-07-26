func minOperations(target, arr []int) int {
    n := len(target)
    pos := make(map[int]int, n)
    for i, val := range target {
        pos[val] = i
    }

    d := []int{}
    // 最长递增子序列
    for _, val := range arr {
        if idx, has := pos[val]; has {  // 过滤不存在于target中的数字
            if p := sort.SearchInts(d, idx); p < len(d) {
                d[p] = idx
            } else {
                d = append(d, idx)
            }
        }
    }
    
    // 总长-最长递增长度
    return n - len(d)
}
