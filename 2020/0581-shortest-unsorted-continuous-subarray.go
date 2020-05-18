// 本题的关键是找出被乱序子数组包含的最小和最大值，
// 找到后只要遍历数组，找到他们原本属于的位置i，j那么就找到了答案
func findUnsortedSubarray(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    var flag1 bool
    var minV int
    for i := 1; i < n; i++ {
        if !flag1 && nums[i] < nums[i-1] {
            flag1 = true
            minV = nums[i]
        }
        if flag1 && nums[i] < minV {
            minV = nums[i]
        }
    }
    // 完全有序
    if !flag1 {
        return 0
    }
    var flag2 bool
    var maxV int
    for i := n - 2; i >= 0; i-- {
        if !flag2 && nums[i] > nums[i+1] {
            flag2 = true
            maxV = nums[i]
        }
        if flag2 && nums[i] > maxV {
            maxV = nums[i]
        }
    }
    left, right := 0, n-1
    // 找到左边界
    for i := 0; i < n; i++ {
        if nums[i] > minV {
            left = i
            break
        }
    }
    // 找到右边界
    for i := n-1; i >= 0; i-- {
        if nums[i] < maxV {
            right = i
            break
        }
    }
    return right - left + 1
}
