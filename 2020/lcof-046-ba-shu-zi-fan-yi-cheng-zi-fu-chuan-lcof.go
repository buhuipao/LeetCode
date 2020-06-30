func translateNum(num int) int {
    nums := []int{}
    for num != 0 {
        nums = append(nums, num % 10)
        num /= 10
    }
    n := len(nums)
    // 初始化
    a, b := 1, 1
    for i := n-2; i >= 0; i-- {
        tmp := 10 * nums[i+1] + nums[i]
        // 两位的数值少于10或者大于25则只有一种解法，否则更新前后两种解法
        if tmp < 10 || tmp > 25 {
            a, b = b, b 
        } else {
            a, b = b, a+b
        }
    }

    return b
}
