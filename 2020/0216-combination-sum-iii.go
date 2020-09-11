func combinationSum3(k int, n int) [][]int {
    // 不存在结果
    if k * 9 < n {
        return nil
    }
    ans := make([][]int, 0)
    backtrack(n, k, 1, []int{}, &ans)
    return ans
}

func backtrack(n, k int, next int, path []int, ans *[][]int) {
    if n == 0 && len(path) == k {
        *ans = append(*ans, path)
        return
    }
    if len(path) == k {
        return
    }
    for i := next; i < 10; i++ {
        // 剪枝
        if n < i {
            return
        }
        // 做出选择
        newP := make([]int, len(path))
        copy(newP, path)
        newP = append(newP, i)
        backtrack(n-i, k, i+1, newP, ans)
        // 撤销选择
    }
}
