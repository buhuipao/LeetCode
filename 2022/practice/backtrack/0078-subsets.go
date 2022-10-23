func subsets(nums []int) [][]int {
    var ans [][]int
    bracktrack(nums, 0, nil, &ans)
    return ans
}

func bracktrack(nums []int, index int, path []int, paths *[][]int) {
    // 添加路径
    *paths = append(*paths, path)

    for i := index; i < len(nums); i++ {
        // 做出选择
        tmp := make([]int, len(path))
        copy(tmp, path)
        tmp = append(tmp, nums[i])

        // 进入下层决策
        bracktrack(nums, i+1, tmp, paths)

        // 撤销选择
    }
}
