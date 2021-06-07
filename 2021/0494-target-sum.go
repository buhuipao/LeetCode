// 递归（回溯）
func findTargetSumWays(nums []int, S int) int {
    if len(nums) == 0 {
        return 0
    }
    ans := 0
    helper(nums, 0, S, &ans)
    return ans
}

func helper(nums []int, i, target int, ans *int) {
    n := len(nums)
    if i == n && target == 0 {
        (*ans)++
        return
    }
    if i < n {
        helper(nums, i+1, target+nums[i], ans)
        helper(nums, i+1, target-nums[i], ans)
    }
}
