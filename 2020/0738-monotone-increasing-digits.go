func monotoneIncreasingDigits(n int) int {
    s := []byte(strconv.Itoa(n))
    i := 1
    // 找到递增的前一部分
    for i < len(s) && s[i] >= s[i-1] {
        i++
    }
    if i < len(s) {
        // 倒退几位，直到满足要求
        for i > 0 && s[i] < s[i-1] {
            s[i-1]--
            i--
        }
        // 补齐9
        for i++; i < len(s); i++ {
            s[i] = '9'
        }
    }
    ans, _ := strconv.Atoi(string(s))
    return ans
}
