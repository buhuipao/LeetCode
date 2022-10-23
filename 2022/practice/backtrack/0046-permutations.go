func permute(nums []int) [][]int {
    var ans [][]int

    backtrack(nums, &ans, []int{}, 0)

    return ans
}

func backtrack(nums []int, paths *[][]int, path []int, index int) {
    // 结束条件
    if index == len(nums) {
        *paths = append(*paths, path)
        return
    }

    for i := index; i < len(nums); i++ {
        // 做出选择
        nums[index], nums[i] = nums[i], nums[index]
        tmp := make([]int, len(path))
        copy(tmp, path)
        tmp = append(tmp, nums[index])
        // 进入下一层做决策
        backtrack(nums, paths, tmp, index+1)
        // 撤销选择
        nums[index], nums[i] = nums[i], nums[index]
    }
}
