func findSubstring(s string, words []string) []int {
    wordL := len(words[0])
    needs := make(map[string]int)
    for i := range words {
        needs[words[i]] += 1
    }

    var left, right, c int
    var ans []int
    window := make(map[string]int, len(needs))
    for i := 0; i < wordL; i++ { // 需要选择不同下标开始
        left, right, c = i, i, 0
        window = make(map[string]int, len(needs))

        for right <= len(s)-wordL {
            w := s[right:right+wordL]
            window[w] += 1
            right += wordL
            c += 1
            if _, ok := needs[w]; !ok { // 杂项
                left, c = right, 0 // 需要跳过一个完整单词
                window = make(map[string]int, len(needs))
                continue
            }

            for window[w] > needs[w] {
                window[s[left:left+wordL]] -= 1
                c -= 1
                left += wordL
            }

            if c == len(words) {
                ans = append(ans, left)
            }
        } 
    }

    return ans
}
