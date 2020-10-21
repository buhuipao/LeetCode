// 双指针
func isLongPressedName(name string, typed string) bool {
    n, m := len(typed), len(name)
    // 特判
    if m == 0 || n == 0 {
        return n == 0 && m == 0
    }
    i, j := 0, 0
    for i < m && j < n {
        // 优先判断相等
        if name[i] == typed[j] {
            i++
            j++
            continue
        } else {
            // 如果不相等，则消除重复字符后再比较
            for i > 0 && j < n && typed[j-1] == typed[j] {
                j++
            }
            if i < m && j < n && name[i] != typed[j] {
                return false
            }
        }
    }
    // typed提前判断完，证明输入不全，返回false
    if i != m {
        return false
    }
    // name遍历完了，需要判断typed之后是否都是末尾的重复字符
    if j != n {
         for j < n && typed[j-1] == typed[j] {
             j++
         }
    }
    return j == n
}
