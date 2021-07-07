// 回溯
func permutation(s string) []string {
    m := make(map[string]interface{})
    helper(s, []byte{}, m)
    ans := make([]string, 0, len(m))
    for k := range m {
        ans = append(ans, k)
    }
    return ans
}

func helper(s string, path []byte, m map[string]interface{}) {
    if s == "" {
        m[string(path)] = nil
        return
    }
    for i := 0; i < len(s); i++ {
        newP := make([]byte, len(path))
        copy(newP, path)
        newP = append(newP, s[i])
        helper(s[:i]+s[i+1:], newP, m)
    }
}
