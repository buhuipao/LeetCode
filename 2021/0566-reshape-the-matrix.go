func matrixReshape(nums [][]int, r int, c int) [][]int {
    rows := len(nums)
    cols := len(nums[0])
    if rows * cols != r * c {
        return nums
    }
    ans := make([][]int, r)
    for i := range ans {
        ans[i] = make([]int, c)
    }
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            tmp := i * cols + j + 1
            I, J := (tmp-1) / c, (tmp-1) % c
            ans[I][J] = nums[i][j]
        }
    }
    return ans
}
