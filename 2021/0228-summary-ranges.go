func summaryRanges(nums []int) []string {
    if len(nums) == 0 {
        return nil
    }
    ans := make([]string, 0)
    ans = append(ans, fmt.Sprintf("%d", nums[0]))
    n := len(nums)
    var flag bool
    for i := 1; i < n; i++ {
        if nums[i] == nums[i-1]+1 {
            flag = true
            continue
        }
        if flag {
            ans[len(ans)-1] = fmt.Sprintf("%s->%d", ans[len(ans)-1], nums[i-1])
            flag = false
        }
        ans = append(ans, fmt.Sprintf("%d", nums[i]))
    }
    // ans最后一个元素：既不包含nums[n-1]，又不以"->nums[n-1]"结尾，那么必须追加到最后
    if ans[len(ans)-1] != fmt.Sprintf("%d", nums[n-1]) && !strings.HasSuffix(ans[len(ans)-1], fmt.Sprintf("->%d", nums[n-1])) {
        ans[len(ans)-1] = fmt.Sprintf("%s->%d", ans[len(ans)-1], nums[n-1])
    }
    return ans
}
