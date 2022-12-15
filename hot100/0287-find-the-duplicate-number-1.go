func findDuplicate(nums []int) int {
    n := len(nums)
    ans := 0
    bit_max := 31
    // 找到数字N的最高位比特位
    for ((n-1) >> bit_max) == 0 {
        bit_max--
    }

    // 
    for bit := 0; bit <= bit_max; bit++ {
        x, y := 0, 0
        // 统计每个比特位1的数量
        for i := 0; i < n; i++ {
            // nums的数字，此bit为1的实际数量
            if (nums[i] & (1 << bit)) > 0 {
                x++
            }
            // 1~N数字中，此bit位为1的预期数量
            if i & (1 << bit) > 0 {
                y++
            }
        }

        // 推导过程为：如果target在此bit位贡献了1
        // 假设缺失的数字此bit位也为1，那么最终x=y+1
        // 假如缺失的数字此bit位为0，则x=y+c，c为target数量
        if x > y {
            ans |= 1 << bit
        }
    }
   
    return ans
}
