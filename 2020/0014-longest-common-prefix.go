// 分治+递归
func longestCommonPrefix(strs []string) string {
    n := len(strs)
    if n == 0 {
        return ""
    }
    if n == 1 {
        return strs[0]
    }
    if n == 2 {
        return helper(strs[0], strs[1])
    }
    return helper(longestCommonPrefix(strs[:n/2]), longestCommonPrefix(strs[n/2:]))
}

func helper(s1, s2 string) string {
    var ans string
    for i := 0; i < len(s1) && i < len(s2); i++ {
        if s1[i] != s2[i] {
            break
        }
        ans = s1[:i+1]
    }
    return ans
}
