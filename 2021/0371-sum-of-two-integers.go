// 使用位操作
// 加法：异或
// 进位：与后左移动
// 直到为零
func getSum(a int, b int) int {
    // 递归退出条件
    if b == 0 {
        return a
    }
    a, b = a ^ b, a & b << 1
    return getSum(a, b)
}
