// kmp记不住
func strStr(haystack string, needle string) int {
    l := len(needle)
    if l == 0 {
        return 0
    }
    n := len(haystack)
    if n < l {
        return -1   // 绝无可能
    }
    indexList := make([]int, 0, n/4)
    for i := l-1; i < n; i++ {      // 先记录所有needle尾巴出现的index，然后去haystack里找找就行
        if haystack[i] == needle[l-1] {
            indexList = append(indexList, i)
        }
    }
    for _, i := range indexList {
        if needle == haystack[i+1-l:i+1] {
            return i+1-l
        }
    }
    return -1
}
