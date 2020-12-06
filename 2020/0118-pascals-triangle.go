func generate(numRows int) [][]int {
    ans := make([][]int, 0, numRows)
    ans = append(ans, [][]int{
        []int{1},
        []int{1, 1},
    }...)
    if numRows < 3 {
        return ans[:numRows]
    }
    for i := 2; i < numRows; i++ {
        tmp := make([]int, 0, i+1)
        tmp = append(tmp, 1)
        for j := 1; j < i; j++ {
            tmp = append(tmp, ans[i-1][j-1]+ans[i-1][j])
        }
        tmp = append(tmp, 1)
        ans = append(ans, tmp)
    }
    return ans
}
