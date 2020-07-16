func searchInsert(nums []int, target int) int {
    L, R := 0, len(nums) - 1
    // 必须要给L越过R的计划，所以退出条件为L>R
    for L <= R {
        M := L + (R - L) / 2
        if nums[M] < target {
            L = M + 1
        } else if nums[M] == target {
            return M
        } else {
            R = M - 1
        }
    }
    return L
}
