// 优化版竖乘法
func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    m, n := len(num1), len(num2)
    ans := make([]byte, m+n)
    var n1, n2 byte
    for i := m-1; i >= 0; i-- {
        n1 = num1[i] - '0'
        for j := n-1; j >= 0; j-- {
            n2 = num2[j] - '0'
            sum := ans[i+j+1] + n1 * n2
            ans[i+j+1] = sum % 10
            ans[i+j] += sum / 10
        }
    }
    // 去除前缀零
    for i := range ans {
        if ans[i] != byte(0) {
            ans = ans[i:]
            break
        }
    }
    var ret bytes.Buffer
    for i := range ans {
        ret.WriteString(strconv.Itoa(int(ans[i])))
    }
    return ret.String()
}
