func sortString(s string) string {
    n := len(s)
    // 统计频率
    m := make([]int, 26)
    for _, c := range s {
        m[int(c-'a')] += 1
    }
    // 写入结果
    ans := make([]byte, 0, len(s))
    for len(ans) < n {
        for i := 0; i < 26; i++ {
            if m[i] > 0 {
                ans = append(ans, 'a' + byte(i))
                m[i] -= 1
            }
        }
        for i := 25; i >= 0; i-- {
            if m[i] > 0 {
                ans = append(ans, 'a' + byte(i))
                m[i] -= 1
            }
        }
    }
    return string(ans)
}
