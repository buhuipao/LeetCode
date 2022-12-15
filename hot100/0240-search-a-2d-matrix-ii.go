// 从左下角或者右上角开始找
// 以右上角为例，分析如下：
// 1）如果target大于它，那就需要排除所在行了往下走，因为它是右上角，它已经是当前行最大值了
// 2）如果target小于它，那就需要排除当前所在列了，因为它是右上角，它已经是当前列最小值了
func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix) 
    if m == 0 {
        return false
    }
    n := len(matrix[0])
    i, j := 0, n-1
    // 查找是否越界
    for i < m && j >= 0 {
        if matrix[i][j] > target {
            j--
        } else if matrix[i][j] < target {
            i++
        } else {
            return true
        }
    } 

    return false
}
