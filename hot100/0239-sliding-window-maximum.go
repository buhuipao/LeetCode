func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 || k == 0 {
        return nil
    }

    var deque []int
    // 初始化
    for i := 0; i < k; i++ {
        cleandeque(&deque, nums, i, k)
        deque = append(deque, i)
    }

    var ans []int
    ans = append(ans, nums[deque[0]])
    for i := k; i < len(nums); i++ {
        cleandeque(&deque, nums, i, k)
        deque = append(deque, i)
        ans = append(ans, nums[deque[0]])
    }

    return ans
}

// cleandeque维护一个单调栈
func cleandeque(deque *[]int, nums []int, i, k int) {
    // 移除即将不在窗口内的索引i-k
    if len(*deque) > 0 && (*deque)[0] == i - k {
        *deque = (*deque)[1:]
    }

    // 从后往前，持续移除deque中nums[index]小于nums[i]的index
    // 直到deque中的index的值（nums[index])大于或者等于nums[i]才停止
    for len(*deque) > 0 && nums[i] > nums[(*deque)[len(*deque)-1]] {
        *deque = (*deque)[:len(*deque)-1]
    }
}
