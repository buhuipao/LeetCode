func minWindow(s string, t string) string {
    need := make(map[byte]int, 0)
    for i := range t {
        need[t[i]] += 1
    }
    cur := make(map[byte]int, 0)
    left, right := -1, len(s)-1
    L, R := 0, -1
    // 记录满足要求的字符串个数
    var valid int
    for i := range s {
        // 右移动窗口
        R++
        // 更新数据
        if _, ok := need[s[i]]; ok {
            cur[s[i]] += 1
            if cur[s[i]] == need[s[i]] {
                valid++
            }
        }
        // 如果当前状态符合要求
        if valid == len(need) {
            // 持续缩紧左边
            for true {
                if R - L < right - left {
                    left, right = L, R
                }
                // 假如存在于目标字符串里
                if need[s[L]] > 0 {
                    cur[s[L]] -= 1
                }
                L++
                // 缩紧左边不满足则推出循环
                if need[s[L-1]] > cur[s[L-1]] {
                    valid--
                    break
                }
            }
        }       
    }
    if left == -1 {
        return ""
    }
    return s[left:right+1]
}
