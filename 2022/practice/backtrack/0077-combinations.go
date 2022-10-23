func combine(n int, k int) [][]int {
    var ans [][]int

    backtrack(n, 1, k, nil, &ans)

    return ans
}

func backtrack(n, cur, k int, path []int, ans *[][]int) {
    if k == 0 {
        tmp := make([]int, len(path))
        copy(tmp, path)
        *ans = append(*ans, tmp)
        return
    }

    for i := cur; i <= n; i++ {
        // 做出选择
        path = append(path, i)

        // 进入下一层决策
        backtrack(n, i+1, k-1, path, ans)

        // 撤销选择
        path = path[:len(path)-1]
    }
}
