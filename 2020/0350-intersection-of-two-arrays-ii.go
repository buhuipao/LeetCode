func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    n1, n2 := len(nums1), len(nums2)
    var i, j int
    ans := make([]int, 0)
    for i < n1 && j < n2 {
        if nums1[i] > nums2[j] {
            j++
        } else if nums1[i] < nums2[j] {
            i++
        } else {
            ans = append(ans, nums1[i])
            i++
            j++
        }
    }
    return ans
}
