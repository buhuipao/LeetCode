// 回溯
func restoreIpAddresses(s string) []string {
    ans := make([]string, 0)
    backtrack(s, 0, 0, []byte{}, &ans)
    return ans
}

func backtrack(s string, dot, index int, path []byte, ans *[]string) {
    // 回溯退出
    if dot == 3 {
        if valid(s[index:]) {
            *ans = append(*ans, string(append(path, []byte(s[index:])...)))
        }
        return
    }
    n := len(s)
    for i := 1; i <= 3; i++ {
        next := i + index
        // 越界
        if next >= n {
            return
        }
        newPath := make([]byte, len(path))
        copy(newPath, path)
        // 剪枝
        if valid(s[index:next]) {
            newPath = append(newPath, []byte(s[index:next])...)
            newPath = append(newPath, '.')
            backtrack(s, dot+1, next, newPath, ans)
        }
    }
}

func valid(s string) bool {
    // 排除前导0
    if s[0] == '0' && len(s) > 1 {
        return false
    }
    // 检查数值不大于255
    if a, err := strconv.Atoi(s); err != nil || a > 255 {
        return false
    }
    return true
}
