func combinationSum(candidates []int, target int) [][]int {
    if len(candidates) == 0 {
        return nil
    }

    sort.Ints(candidates)
    var ans [][]int
    backtrack(candidates, nil, 0, target, &ans)
    return ans
}

func backtrack(nums []int, path []int, index, target int, ans *[][]int) {
    // 退出条件
    if target == 0 {
        tmp := make([]int, len(path))
        copy(tmp, path)
        *ans = append(*ans, tmp)
    }

    for i := index; i < len(nums); i++ {
        // 剪枝
        if nums[i] > target {
            return
        }

        // 做出选择
        path = append(path, nums[i])
        // 递归（可以重复取，所以依然是从i开始）
        backtrack(nums, path, i, target-nums[i], ans)
        // 撤销选择
        path = path[:len(path)-1]
    }
}
