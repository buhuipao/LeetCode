// 回溯
func findSubsequences(nums []int) [][]int {
    ans := make([][]int, 0)
    backtrack(nums, 0, []int{}, &ans)
    return ans
}

func backtrack(nums []int, index int, path []int, ans *[][]int) {
    if len(path) >= 2 {
        *ans = append(*ans, path)
    }
    m := make(map[int]bool, 200)
    for i := index; i < len(nums); i++ {
        if len(path) != 0 && path[len(path)-1] > nums[i] {
            continue
        }
        // 剪枝
        if m[nums[i]] {
            continue
        }
        newPath := make([]int, len(path))
        copy(newPath, path)
        newPath = append(newPath, nums[i])
        m[nums[i]] = true

        backtrack(nums, i+1, newPath, ans)

    }
}
