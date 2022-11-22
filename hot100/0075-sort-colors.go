// 三向切分快排，使得：
// 1) nums[0:left]都小于1；
// 2）nums[left:i]都等于1；
// 3）nums[i:right]待定;
// 4）nums[right:n]都大于1；
func sortColors(nums []int)  {
    left, i, right := 0, 0, len(nums)
    for i < right {
        switch {
            case nums[i] < 1:
                nums[left], nums[i] = nums[i], nums[left]
                left++
                i++
            
            case nums[i] == 1:
                i++
            
            default:
                // 替换right前面一个
                nums[i], nums[right-1] = nums[right-1], nums[i]
                right--
        }
    }
}
