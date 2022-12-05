// DP:
// maxV[i] = max(maxV[i-1]*nums[i], nums[i])
// minV[i] = max(minV[i-1]*nums[i], nums[i])
func maxProduct(nums []int) int {
    ans := ^int(^uint(0)>>1)
    maxV, minV := 1, 1
    for i := range nums {
        if nums[i] < 0 {
            maxV, minV = minV, maxV
        }

        maxV = max(maxV*nums[i], nums[i])
        minV = min(minV*nums[i], nums[i])
        ans = max(maxV, ans)
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    
    return b
}
