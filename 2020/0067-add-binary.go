func addBinary(a string, b string) string {
    i, j := len(a)-1, len(b)-1
    if i < j {
        return addBinary(b, a)
    }
    var x, y byte
    // 默认进位值为'0'
    y = '0'
    // 给最终的sum留下进位空间
    ans := make([]byte, i+2)
    for i >= 0 || j >= 0 || y != '0' {
        var b1, b2 byte
        b1, b2 = '0', '0'
        if i >= 0 {
            b1 = a[i]
        }
        if j >= 0 {
            b2 = b[j]
        }
        x, y = add(b1, b2, y)
        ans[i+1] = x
        i--
        j--
    }
    // 去除未进位的前缀0
    k := 0
    for ans[k] == byte(0) {
        k++
    }
    return string(ans[k:])
}

// 求和计算
func add(a, b, y byte) (byte, byte) {
    sum := (a - '0') + (b - '0') + (y - '0')
    return byte(sum % 2) + '0', byte(sum / 2) + '0'
}
