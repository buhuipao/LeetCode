func deleteAndEarn(nums []int) (ans int) {
    sort.Ints(nums)
    sum := []int{nums[0]}
    // 把nums分组，把相邻的数字放在一起，这样分组后的各组之间无影响
    for i := 1; i < len(nums); i++ {
        if val := nums[i]; val == nums[i-1] {
            sum[len(sum)-1] += val
        } else if val == nums[i-1]+1 {
            sum = append(sum, val)
        } else {
            ans += rob(sum)
            sum = []int{val}
        }
    }
    ans += rob(sum)
    return
}

func rob(nums []int) int {  // 打家劫舍
    if len(nums) == 1 {
        return nums[0]
    }
    first, second := nums[0], max(nums[0], nums[1])
    for i := 2; i < len(nums); i++ {
        first, second = second, max(first+nums[i], second)
    }
    return second
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

