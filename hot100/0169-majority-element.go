/*
我们维护一个计数器，如果遇到一个我们目前的候选众数，就将计数器加一，否则减一。
只要计数器等于0 ，我们就将nums中之前访问的数字全部忘记，并把下一个数字当做候选的众数。
*/
// 投票算法
func majorityElement(nums []int) int {
    if len(nums) == 0 {
        return -1
    }

    ans, count := nums[0], 1
    for i := 1; i < len(nums); i++ {
        if count == 0 {
            ans = nums[i]
            count++
            continue
        }

        if nums[i] == ans {
            count++
        } else {
            count--
        }
    }

    return ans
}
