// Leetcode 不会系列
func numberOfArithmeticSlices(nums []int) (ans int) {
    f := make([]map[int]int, len(nums))
    for i, x := range nums {
        f[i] = map[int]int{}
        for j, y := range nums[:i] {
            d := x - y
            cnt := f[j][d]
            ans += cnt
            f[i][d] += cnt + 1
        }
    }
    return
}
