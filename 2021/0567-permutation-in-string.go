func checkInclusion(s1 string, s2 string) bool {
    // 统计S1的词频
    m, m1 := make(map[byte]int), make(map[byte]bool)
    for i := range s1 {
        m[s1[i]] += 1
        m1[s1[i]] = true
    }
    // 滑动窗口
    L := 0
    for i := range s2 {
        // s1中不存在的字符串时重置L，同时重新计数此段跳过的字符
        if !m1[s2[i]] {
            for L < i {
                if m1[s2[L]] {
                    m[s2[L]] += 1
                }
                L++
            }
            L = i + 1
            continue
        }
        m[s2[i]] -= 1
        if m[s2[i]] == 0 {
            delete(m, s2[i])
        }
        // 如果字符数量超过所需要的，需要移动L直到符合要求
        if m[s2[i]] < 0 {
            for m[s2[i]] < 0 {
                m[s2[L]] += 1
                L++
            }
            delete(m, s2[i])    // 注意删除字符
        }
        if len(m) == 0 {
            return true
        }
    }
    return false
}
