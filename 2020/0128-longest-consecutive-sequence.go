func longestConsecutive(nums []int) int {
    m := make(map[int]bool, len(nums))
    for i := range nums {
        m[nums[i]] = true
    }
    ans := 0
    for k := range m {
        // 如果有比当前值更小的，就以更小值为准
        if _, ok := m[k-1]; ok {
            continue
        }
        tmp := 1
        for m[k+1] {
            tmp++
            k++
        }
        if tmp > ans {
            ans = tmp
        }
    }
    return ans
}
