func subarraySum(nums []int, k int) int {
    m := make(map[int]int)
    var ans, sum int
    for i := range nums {
        sum += nums[i]
        if sum == k {
            ans++
        }
        ans += m[sum-k]
        m[sum] +=1
    }

    return ans
}
