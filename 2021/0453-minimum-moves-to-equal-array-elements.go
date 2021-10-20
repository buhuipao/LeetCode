// 思路一：数学问题：sum + x * (n-1) = n * L
// (n-1) * (L-x) = sum - L 求解方程即可

// 思路二：n-1个数加一，等效于1个数减一
// 所以只要找到最小值，然后统计所有数字变成它需要的步数
func minMoves(nums []int) (ans int) {
    min := nums[0]
    for _, num := range nums[1:] {
        if num < min {
            min = num
        }
    }

    for _, num := range nums {
        ans += num - min
    }
    return
}
