func findUnsortedSubarray(nums []int) int {
    L := -1
    minV := int(^uint(0)>>1)
    for i := len(nums)-1; i >= 0; i-- {
        // 如果当前值更小，更新遍历过的nums[i]获取到的最小值minV
        if nums[i] <= minV {
            minV = nums[i]
        // 如果当前值大于最小值，证明当前索引i很可能是待排序左边界L
        } else {
            L = i
        }
    }

    R := -1
    maxV := ^int(^uint(0)>>1) // 故意取一个MinInt
    for i := range nums {
        // 如果当前值更大，则更新最大值
        if nums[i] >= maxV {
            maxV = nums[i]
        // 如果当前值不按照常理，反而小于前面遍历获得的最大值，证明此处索引i应该为待排序边界R
        } else {
            R = i
        }
    }

    if L == -1 || R == -1 {
        return 0
    }

    return R - L + 1
}
