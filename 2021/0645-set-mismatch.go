func findErrorNums(nums []int) []int {
    xor := 0
    for _, v := range nums {
        xor ^= v
    }
    n := len(nums)
    for i := 1; i <= n; i++ {
        xor ^= i
    }

    // 找到两个数字不相同的最低bit位
    lowbit := xor & -xor
    num1, num2 := 0, 0
    // 利用最低bit位区分阵营
    // 因为左边的target1重复了，异或后和右边的target2缺失是一样的效果
    for _, v := range nums {
        if v&lowbit == 0 {
            num1 ^= v
        } else {
            num2 ^= v
        }
    }
    // 利用异或满足交换律计算得出target1、target2
    for i := 1; i <= n; i++ {
        if i&lowbit == 0 {
            num1 ^= i
        } else {
            num2 ^= i
        }
    }
    // 找出重复的数
    for _, v := range nums {
        if v == num1 {
            return []int{num1, num2}
        }
    }
    
    return []int{num2, num1}
}
