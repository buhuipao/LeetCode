// 递归·回溯
func combinationSum(candidates []int, target int) [][]int {
    if len(candidates) == 0 {
        return nil
    }
    sort.Ints(candidates)
    ans := make([][]int, 0)
    path := make([]int, 0)
    backtrack(candidates, 0, target, path, &ans)
    return ans
}

func backtrack(nums []int, index int, target int, path []int, ans *[][]int) {
    // 终止
    if target == 0 {
        *ans = append(*ans, path)
        return
    }
    n := len(nums)
    for i := index; i < n; i++ {
        // 剪枝，如果目标值已经小于最小待选择的数字了就没必要再进行回溯
        if target-nums[i] < 0 {
            return
        }
        // 选择
        newP := make([]int, len(path))
        copy(newP, path)
        newP = append(newP, nums[i])
        // 递归
        backtrack(nums, i, target-nums[i], newP, ans)
        // 撤销选择
        //newP = newP[:len(newP)-1]
    }
}
