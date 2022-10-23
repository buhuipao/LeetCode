func subsetsWithDup(nums []int) [][]int {
    var ans [][]int
    sort.Ints(nums)
    bracktrack(nums, 0, nil, &ans)
    return ans
}

func bracktrack(nums []int, index int, path []int, paths *[][]int) {
    // 添加路径
    tmp := make([]int, len(path))
    copy(tmp, path)
    *paths = append(*paths, tmp)

    for i := index; i < len(nums); i++ {
        // 做出选择
        path = append(path, nums[i])

        // 进入下层决策
        bracktrack(nums, i+1, path, paths)

        // 撤销选择
        path = path[:len(path)-1]

        // 跳过重复项
        for i < len(nums)-1 && nums[i] == nums[i+1] {
            i++
        }
    }
}
