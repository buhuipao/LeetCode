// 两遍扫描，找到数列中倒数第一个升序和降序，然后替换降序中
// [1, 2, 3, 5, 4, 2] => [1, 2, 4,  2, 3, 5] 
func nextPermutation(nums []int)  {
    n := len(nums)
    if n < 2 {
        return
    }

    // 1. 从后往前找到第一个升序数组对
    i := n - 2
    for ; i >= 0; i-- {
        if nums[i] < nums[i+1] {
            break
        }
    }

    // 2. 将nums[i]与其后的降序数组中一个稍大于它的数字进行一次互换
    // [1, 2, 2, 4, 3, 1] 将nums[2]与3进行替换
    if i >= 0 {
        for j := n-1; j > i; j-- {
            if nums[j] > nums[i] {
                nums[i], nums[j] = nums[j], nums[i]
                break
            }
        }
    }

    // 3. 由第一步可知，nums[i+1:]一定是一个降序数组，反转这个降序数组就是最后的答案
    // i 为 -1时也包括了
    l, r := i+1, n-1
    for l < r {
        nums[l], nums[r] = nums[r], nums[l]
        l++
        r-- 
    }
}