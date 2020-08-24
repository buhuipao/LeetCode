// 采用双拼字符的方法：
// s = x + x + n * x，n >= 0，令 S = s + s = x + x + 2n*x + x + x
// 在对S掐头去尾(S[1:-1])破坏首位两个x的完整性后，S1 = x[1:] + x + 2n * x + x + x[:-1]
// 如果s为重复字符串，那么S1中至少可以找到一个完整的s，若找不到则s不是重复字符串
func repeatedSubstringPattern(s string) bool {
    if s == "" {
        return true
    }
    S, n := s + s, len(s)
    for i := 1; i < n; i++ {
        if S[i:i+n] == s {
            return true
        }
    } 
    return false
}
