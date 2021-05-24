// 类似于前缀和
func kthLargestValue(matrix [][]int, k int) int {
    nums := make([]int, 0, len(matrix)*len(matrix[0]))
    for i := range matrix {
        for j := range matrix[i] {
            var v1, v2, v3 int
            if i >= 1 && j >= 1 {
                v1 = matrix[i-1][j-1]
                v2 = matrix[i][j-1]
                v3 = matrix[i-1][j]
            } else if i >= 1 {
                v1 = matrix[i-1][0]
            } else if j >= 1 {
                v2 = matrix[0][j-1]
            }
            matrix[i][j] ^= v1 ^ v2 ^ v3
            // 此处可以优化为大顶堆
            nums = append(nums, matrix[i][j])
        }
    }
    sort.Ints(nums)
    return nums[len(nums)-k]
}
