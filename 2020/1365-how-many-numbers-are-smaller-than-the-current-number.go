func smallerNumbersThanCurrent(nums []int) []int {
    n := len(nums)
    tmp := make([]int, len(nums))
    copy(tmp, nums)
    sort.Ints(tmp)

    m := make(map[int]int)
    for i := 1; i < n; i++ {
        if tmp[i] == tmp[i-1] {
            m[tmp[i]] = m[tmp[i-1]]
        } else {
            m[tmp[i]] = i
        }
    }
    for i := range nums {
        tmp[i] = m[nums[i]]
    }
    return tmp
}
