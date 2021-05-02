// 统计32个bit位1的计数
func singleNumber(nums []int) int {
    ans := int32(0)
    for i := 0; i < 32; i++ {
        total := int32(0)
        for _, num := range nums {
            total += int32(num) >> i & 1
        }
        if total%3 > 0 {    // 多出来的位就是目标数的
            ans |= 1 << i
        }
    }
    return int(ans)
}
