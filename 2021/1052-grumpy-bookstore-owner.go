// 满意度 = max(在X窗口内由于使用技能增加的满意度) + 原本的满意度
func maxSatisfied(customers []int, grumpy []int, X int) int {
    total := 0
    n := len(customers)
    for i := 0; i < n; i++ {
        if grumpy[i] == 0 {
            total += customers[i]
        }
    }
    increase := 0
    for i := 0; i < X; i++ {
        increase += customers[i] * grumpy[i]
    }
    maxIncrease := increase
    for i := X; i < n; i++ {
        increase = increase - customers[i-X]*grumpy[i-X] + customers[i]*grumpy[i]
        maxIncrease = max(maxIncrease, increase)
    }
    return total + maxIncrease
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
