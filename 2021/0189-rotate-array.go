// 三次翻转
func rotate(nums []int, k int)  {
    n := len(nums)
    if n <= 1 {
        return 
    }
    k = k % n
    // 1、整体翻转
    for i := 0; i < n/2; i++ {
        nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
    }
    // 2、翻转点原右边翻转
    for i := 0; i < k/2; i++ {
        nums[i], nums[k-1-i] = nums[k-1-i], nums[i]
    } 
    // 3、翻转点原左边翻转
    for i := k; i < (n+k)/2; i++ {
        nums[i], nums[n+k-1-i] = nums[n+k-1-i], nums[i]
    }
}
