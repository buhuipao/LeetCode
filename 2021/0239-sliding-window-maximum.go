func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 || k == 0 {
        return nil
    }
    ret := []int{}
    // 初始化
    deq := []int{}
    for i := 0; i < k; i++ {
        cleanDeque(&deq, nums, i, k)
        deq = append(deq, i)
    }
    ret = append(ret, nums[deq[0]])
    for i := k; i < len(nums); i++ {
        cleanDeque(&deq, nums, i, k)
        deq = append(deq, i)
        ret = append(ret, nums[deq[0]])
    }
    return ret
}

func cleanDeque(deq *[]int, nums []int, i int, k int) {
    // 移除不在窗口内的索引
    if len(*deq) != 0 && (*deq)[0] == i - k {
        *deq = (*deq)[1:]
    }
    // 移除所有比nums[i]小的索引号
    for len(*deq) != 0 && nums[i] > nums[(*deq)[len(*deq)-1]] {
        *deq = (*deq)[:len(*deq)-1]
    }
}
