func longestConsecutive(nums []int) int {
    m := make(map[int]bool, len(nums))
    for i := range nums {
        m[nums[i]] = true
    }

    var ans int
    for k := range m {
        // 如果存在nums[i]-1，则期待往下找更小的值 
        if _, ok := m[k-1]; ok {
            continue
        }

        tmp := 0
        for m[k] {
            tmp++
            k++
        }

        if tmp > ans {
            ans = tmp
        }
    }

    return ans
}
