// 1、先计算优先级高的运算并用stack保存
// 2、再把stack的数字按加法计算一遍
func calculate(s string) (ans int) {
    stack := []int{}
    preSign := '+'
    num := 0
    for i, ch := range s {
        isDigit := '0' <= ch && ch <= '9'
        if isDigit {
            num = num*10 + int(ch-'0')
        }
        if !isDigit && ch != ' ' || i == len(s)-1 {
            switch preSign {
            case '+':
                stack = append(stack, num)
            case '-':
                stack = append(stack, -num)
            case '*':
                stack[len(stack)-1] *= num
            default:
                stack[len(stack)-1] /= num
            }
            preSign = ch
            num = 0
        }
    }
    for _, v := range stack {
        ans += v
    }
    return
}
