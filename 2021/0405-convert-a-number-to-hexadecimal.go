// 32位整数可以看出8个4位16进制数
func toHex(num int) string {
    if num == 0 {
        return "0"
    }
    sb := &strings.Builder{}
    for i := 7; i >= 0; i-- {
        val := num >> (4 * i) & 0xf
        if val > 0 || sb.Len() > 0 {
            var digit byte
            if val < 10 {
                digit = '0' + byte(val)
            } else {
                digit = 'a' + byte(val-10)
            }
            sb.WriteByte(digit)
        }
    }
    return sb.String()
}
