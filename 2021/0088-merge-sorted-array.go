func merge(nums1 []int, m int, nums2 []int, n int)  {
    m--
    n--
    for m >= 0 || n >= 0 {
        if n < 0 {
            return
        }
        if m < 0 {
            nums1[n] = nums2[n]
            n--
            continue
        }
        if nums1[m] > nums2[n] {
            nums1[m+n+1] = nums1[m]
            m--
        } else {
            nums1[m+n+1] = nums2[n]
            n--
        }
    }
}
