func firstUniqChar(s string) int {
    m := make([]int, 26)
    for i, c := range s {
        index := int(c-'a')
        if m[index] == 0 {
            m[index] = i+1 // 避免字符不出现，给i加1
        } else { // 重复的
            m[index] = -1
        }
    }
    for _, c := range s {
        index := int(c-'a')
        if m[index] > 0 {
            return m[index]-1
        } 
    }
    return -1
}
