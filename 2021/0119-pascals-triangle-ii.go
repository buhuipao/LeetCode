// C[i] = C1[i] + C1[i-1]
// 倒着计算，这样就只需要一个数组
func getRow(rowIndex int) []int {
    arr := make([]int, rowIndex+1)
    arr[0] = 1
    for i := 1; i <= rowIndex; i++ {
        for j := i; j > 0; j-- {
            arr[j] += arr[j-1]
        }
    }
    return arr
}
