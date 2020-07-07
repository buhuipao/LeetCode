// 滑动窗口
// [shorter1, shorter2, ..., longerk-1, longerk]
func divingBoard(shorter int, longer int, k int) []int {
    if shorter > longer {
        return divingBoard(longer, shorter, k)
    }
    // k为0
    if k == 0 {
        return nil
    }
    // 相同
    if shorter == longer {
        return []int{k * shorter}
    }
    ans := []int{}
    a, b := longer * k, longer - shorter
    for i := k; i >= 0; i-- {
        ans = append(ans, a - i * b)
    }
    return ans
}
