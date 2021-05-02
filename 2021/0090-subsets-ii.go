// 回溯或者递归
func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    ans := make([][]int, 0)
    backtrack(nums, 0, &ans, []int{})
    return ans
}

func backtrack(nums []int, index int, ans *[][]int, path []int) {
    *ans = append(*ans, path)
    if index == len(nums) {
        return 
    }
    for i := index; i < len(nums); i++ {
        // 剪枝
        if i > index && nums[i] == nums[i-1] {
            continue
        }
        // 作出选择
        newP := make([]int, len(path)) 
        copy(newP, path)
        newP = append(newP, nums[i])
        backtrack(nums, i+1, ans, newP)
        // 撤销
    }
}
