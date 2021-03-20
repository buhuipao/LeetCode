// 比较简单, 利用栈的思想, 遍历输入的数组, 依次入栈
// 遇到运算符就出栈两个元素(第一个出栈元素放运算符右边)进行运算,并将结果入栈
// 最后栈里只会剩下计算结果
func evalRPN(tokens []string) int {
    if len(tokens) == 0 {
        return 0
    }
    stack := []int{}
    m := map[string]func(x, y int)int{
        "+": add,
        "-": sub,
        "/": div,
        "*": mult,
    }
    for i := range tokens {
        switch tokens[i] {
            case "+","-","/","*":
                x, y := stack[len(stack)-2], stack[len(stack)-1]
                stack = stack[:len(stack)-2]
                z := m[tokens[i]](x, y)
                stack = append(stack, z)    
            default:
                x, _ := strconv.Atoi(tokens[i])
                stack = append(stack, x)
        }
    }
    return stack[0]
}

func add(x, y int) int {
    return x + y
}

func sub(x, y int) int {
    return x - y
}

func div(x, y int) int {
    return x / y
}

func mult(x, y int) int {
    return x * y
}
