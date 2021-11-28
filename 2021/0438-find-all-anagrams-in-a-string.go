// 滑动窗口
func findAnagrams(s string, p string) []int {
    need := make(map[byte]int)
    for i := range p {
        need[p[i]] += 1
    }
    left, right := 0, -1
    ans := []int{}
    cur := make(map[byte]int, 0)
    count := 0
    for i := range s {
        // 向右扩充
        right++
        // 数据处理
        if _, ok := need[s[i]]; ok {
            cur[s[i]] += 1
            // 刚好符合条件
            if cur[s[i]] == need[s[i]] && count < len(need) {
                count++
            }
        }
        // 缩紧左边
        for count >= len(need) {
            if len(p) == right - left + 1 {
                ans = append(ans, left)
            }
            if need[s[left]] > 0 {

                if cur[s[left]] > 0 {
                    cur[s[left]] -= 1
                }
                if cur[s[left]] < need[s[left]] {
                    count--
                }
            }
            left++
        }
    }
    return ans
}
