// Hash
// a + b + c = d 转化为：d - c = a + b
func countQuadruplets(nums []int) (ans int) {
    cnt := map[int]int{}
    for b := len(nums) - 3; b >= 1; b-- {
        for _, dValue := range nums[b+2:] {
            // 每次都按住c，假设c在b右边一个位置，统计d-c的值
            cValue := nums[b+1]
            cnt[dValue-cValue]++
        }

        for _, aValue := range nums[:b] {
            // sum = a + b
            if sum := aValue + nums[b]; cnt[sum] > 0 {
                ans += cnt[sum]
            }
        }
    }
    return
}
