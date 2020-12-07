/* 分三步走
第一步：将首列全部置位1，保证最高位全部取到，首列不为1的行全部翻转
第二步：从第二列开始，将所有列中1的数量小于0的数量的行翻转，保证取1的数量尽可能多
第三步：计算结果返回
*/
func matrixScore(a [][]int) int {
    m, n := len(a), len(a[0])
    // 计算第一列的数值
    ans := 1 << (n - 1) * m
    for j := 1; j < n; j++ {
        // 统计每一列1的数量
        ones := 0
        for _, row := range a {
            if row[j] == row[0] {
                ones++
            }
        }
        // 是的1的数量大于0
        if ones < m-ones {
            ones = m - ones
        }
        // 累加列的值
        ans += 1 << (n - 1 - j) * ones
    }
    return ans
}
