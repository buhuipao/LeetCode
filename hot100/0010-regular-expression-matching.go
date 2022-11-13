func isMatch(s string, p string) bool {
    // 考虑s或者p为空的情况
    if p == "" {
        return s == ""
    }

    // 先考虑第一个字符的匹配情况
    var match bool
    if len(s) > 0 && (p[0] == '.' || s[0] == p[0]) {
        match = true
    }

    // 考虑两种特殊的带*情况
    // 1）*号代表了0次重复前一字符，比如：s="abc", p="b*abc"
    // 2）*号代表了1+次重复前一字符，比如：s="bbabc", p="b*abc" 或者 p=".*abc"
    if len(p) >= 2 && p[1] == '*' {
        return isMatch(s, p[2:]) || (match && isMatch(s[1:], p))
    }

    // 普通情况
    return match && isMatch(s[1:], p[1:])
}
