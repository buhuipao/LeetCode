func permute(nums []int) [][]int {
    var ans [][]int
    backtrack(nums, nil, 0, &ans)
    return ans
}

func backtrack(nums []int, path []int, index int, ans *[][]int) {
    // 结束条件
    if index == len(nums) {
        tmp := make([]int, len(path))
        copy(tmp, path)
        *ans = append(*ans, tmp)
    }

    for i := index; i < len(nums); i++ {
        // 做出选择
        nums[i], nums[index] = nums[index], nums[i]
        path = append(path, nums[index])

        backtrack(nums, path, index+1, ans)

        // 撤销选择
        path = path[:len(path)-1]
        nums[i], nums[index] = nums[index], nums[i]
    }

}
