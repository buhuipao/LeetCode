func permuteUnique(nums []int) [][]int {
    if len(nums) == 0 {
        return nil
    }
    sort.Ints(nums)
    ans := make([][]int, 0)
    path := make([]int, 0)
    used := make([]bool, len(nums))
    backtrack(path, nums, used, &ans)
    return ans
}

func backtrack(path []int, options []int, used []bool, ans *[][]int) {
    // 终止条件
    if len(path) == len(options) {
        *ans = append(*ans, path)
        return
    }
    n := len(options)
    for i := 0; i < n; i++ {
        // 如果上层用过则直接跳过
        if used[i] {
            continue
        }
        // 剪枝，避免本层中用到前一个重复的数字
        if i > 0 && options[i] == options[i-1] && !used[i-1] {
            continue
        }
        newP := make([]int, len(path))
        copy(newP, path)
        newP = append(newP, options[i])
        // 作出选择
        used[i] = true
        // 递归
        backtrack(newP, options, used, ans)
        // 撤销选择，以便重新作出选择
        used[i] = false
    }
}
