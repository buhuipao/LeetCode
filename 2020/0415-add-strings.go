// 参照链表相加
func addStrings(num1 string, num2 string) string {
    l1, l2 := len(num1), len(num2)
    var ans []byte
    var index int
    if l1 > l2 {
        ans = make([]byte, l1+1)
        index = l1
    } else {
        ans = make([]byte, l2+1)
        index = l2
    }
    var v1, v2 byte
    var carry byte
    carry = '0'
    l1--
    l2--
    for l1 >= 0 || l2 >= 0 || carry != '0' {
        v1, v2 = '0', '0'
        if l1 >= 0 {
            v1 = num1[l1]
            l1--
        }
        if l2 >= 0 {
            v2 = num2[l2]
            l2--
        }
        c, sum := add(v1, v2, carry)
        carry = c
        ans[index] = sum
        index--
    }
    if ans[0] == byte(0) {
        return string(ans[1:])
    }
    return string([]byte(ans))
}

func add(a, b, c byte) (byte, byte) {
    sum := (a - '0') + (b - '0') + (c - '0')
    if sum >= byte(10) {
        return '1', sum - byte(10) + '0'
    } else {
        return '0', sum + '0'
    }
}
