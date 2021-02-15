func findMaxConsecutiveOnes(nums []int) int {
    var ans, tmp int
    for i := range nums {
        if nums[i] == 1 {
            tmp++
        } else {
            tmp = 0
        }
        if tmp > ans {
            ans = tmp
        }
    }
    return ans
}
