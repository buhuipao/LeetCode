func minAbsoluteSumDiff(nums1, nums2 []int) int {
    rec := append(sort.IntSlice(nil), nums1...)
    rec.Sort()
    sum, maxn, n := 0, 0, len(nums1)
    for i, v := range nums2 {
        diff := abs(nums1[i] - v)
        sum += diff
        j := rec.Search(v)
        if j < n {
            maxn = max(maxn, diff-(rec[j]-v))
        }
        if j > 0 {
            maxn = max(maxn, diff-(v-rec[j-1]))
        }
    }
    return (sum - maxn) % (1e9 + 7)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
