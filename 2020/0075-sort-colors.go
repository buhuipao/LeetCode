// 三向切分快排，使得：
// 1) nums[0:left]都小于1；
// 2）nums[right+1:n]都大于1；
// 3）nums[left:i]都等于1；
// 4）nums[i:right]待定;
func sortColors(nums []int)  {
    left, i, right := 0, 0, len(nums) - 1
    for i <= right {
        switch {
            case nums[i] < 1:
                nums[i], nums[left] = nums[left], nums[i]
                left++
                i++
            case nums[i] == 1:
                i++
            case nums[i] > 1:
                nums[right], nums[i] = nums[i], nums[right]
                right--
        }
    }
}
