func threeSum(nums []int) [][]int {
    n := len(nums)
    if n < 3 {
        return nil
    }
    sort.Ints(nums)
    ans := make([][]int, 0)
    for i := 0; i < n; i++ {
        if nums[i] > 0 {
            break
        }
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        l, r := i+1, n-1
        for l < r {
            value := nums[i] + nums[l] + nums[r]
            if value < 0 {
                l++
            } else if value == 0 {
                ans = append(ans, []int{nums[i], nums[l], nums[r]})
                for l < r && nums[l] == nums[l+1] {
                    l++
                }
                for l < r && nums[r] == nums[r-1] {
                    r--
                }
                l++
                r--
            } else {
                r-- 
            }
        }
    }
    return ans
}
