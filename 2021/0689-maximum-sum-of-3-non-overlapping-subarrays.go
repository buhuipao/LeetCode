// 滑动窗口加强版
// 关键在于通过初始位置，设置固定k间隔，使得三组数字不会出现重叠；然后利用滑动窗口，沿途统计最大值
func maxSumOfThreeSubarrays(nums []int, k int) (ans []int) {
    var sum1, maxSum1, maxSum1Idx int
    var sum2, maxSum12, maxSum12Idx1, maxSum12Idx2 int
    var sum3, maxTotal int
    for i := k * 2; i < len(nums); i++ {
        sum1 += nums[i-k*2]
        sum2 += nums[i-k]
        sum3 += nums[i]
        if i >= k*3-1 {
            if sum1 > maxSum1 {
                maxSum1 = sum1
                maxSum1Idx = i - k*3 + 1
            }
            if maxSum1+sum2 > maxSum12 {
                maxSum12 = maxSum1 + sum2
                maxSum12Idx1, maxSum12Idx2 = maxSum1Idx, i-k*2+1
            }
            if maxSum12+sum3 > maxTotal {
                maxTotal = maxSum12 + sum3
                ans = []int{maxSum12Idx1, maxSum12Idx2, i - k + 1}
            }
            sum1 -= nums[i-k*3+1]
            sum2 -= nums[i-k*2+1]
            sum3 -= nums[i-k+1]
        }
    }
    return
}
