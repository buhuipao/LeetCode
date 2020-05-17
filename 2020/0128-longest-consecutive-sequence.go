// Hash表 或者 并查集
func longestConsecutive(nums []int) int {
    // 特判
    n := len(nums)
    if n == 0 {
        return 0
    }
    // 记录所有值
    record := make(map[int]bool, n)
    for i := range nums {
        record[nums[i]] = true
    }
    ans := 0
    MAX_INT, MIN_INT := int(^uint(0)>>1), ^int(^uint(0)>>1)
    for i := range nums {
        // 如果nums[i]-1也在记录里，那么nums[i]一定会被串起啦，所以不需要参与下面的计算
        if nums[i] != MIN_INT {
            if _, ok := record[nums[i]-1]; ok {
                continue
            }
        }
        // 统计连续序列
        cur := 1
        curV := nums[i]
        for curV != MAX_INT {
            curV++
            if _, ok := record[curV]; ok {
                cur++
            } else {
                break
            }
        }
        if cur > ans {
            ans = cur
        }
    }

    return ans
}
