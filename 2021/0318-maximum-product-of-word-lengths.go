// 1）位运算，把字符串中字母映射为一个int类型数字的bit位
// 2）使用hash表节省循环次数
func maxProduct(words []string) (ans int) {
    masks := map[int]int{}
    for _, word := range words {
        mask := 0
        for _, ch := range word {
            mask |= 1 << (ch - 'a')
        }
        if len(word) > masks[mask] {
            masks[mask] = len(word)
        }
    }

    for x, lenX := range masks {
        for y, lenY := range masks {
            if x&y == 0 && lenX*lenY > ans {
                ans = lenX * lenY
            }
        }
    }
    return
}
