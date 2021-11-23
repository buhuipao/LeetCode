// 分情况
// 1、长度是否相等；
// 2、如果字符串相同就看是否有相同字符
// 3、是否只有两个不同字符，且互换后相等；
func buddyStrings(s, goal string) bool {
    if len(s) != len(goal) {
        return false
    }

    if s == goal {
        seen := [26]bool{}
        for _, ch := range s {
            if seen[ch-'a'] {
                return true
            }
            seen[ch-'a'] = true
        }
        return false
    }

    first, second := -1, -1
    for i := range s {
        if s[i] != goal[i] {
            if first == -1 {
                first = i
            } else if second == -1 {
                second = i
            } else {
                return false
            }
        }
    }
    return second != -1 && s[first] == goal[second] && s[second] == goal[first]
}
