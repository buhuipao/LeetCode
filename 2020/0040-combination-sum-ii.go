func combinationSum2(candidates []int, target int) [][]int {
    ans := make([][]int, 0)
    path := make([]int, 0)
    sort.Ints(candidates)
    backtrack(candidates, target, 0, path, &ans)
    return ans
}


func backtrack(candidates []int, target, index int, path []int, ans *[][]int) {
    n := len(candidates)
    if target == 0 {
        *ans = append(*ans, path)
    }
    if index == n {
        return
    }
    for i := index; i < n; i++ {
        // 跳过相同的数
        rawI := i
        for i < n-1 && candidates[i] == candidates[i+1] {
            i++
        }
        // 剪枝
        if target < candidates[i] {
            return
        }
        // 做出选择
        newP := make([]int, len(path))
        copy(newP, path)
        newP = append(newP, candidates[i])
        backtrack(candidates, target-candidates[i], rawI+1, newP, ans)
        // 撤销选择
    }
}
