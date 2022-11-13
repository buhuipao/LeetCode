// 二分法 + 递归
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m, n := len(nums1), len(nums2)
    if m + n == 0 {
        return 0.0
    }

    // 通过两次查找，合并了奇数和偶数的情况
    k1 := (m + n + 1) / 2
    k2 := (m + n + 2) / 2
    value1 := findK(nums1, nums2, k1)
    value2 := findK(nums1, nums2, k2)

    return float64(value1 + value2) / 2
}

func findK(nums1, nums2 []int, k int) int {
    m, n := len(nums1), len(nums2)
    // 假定num1长度小于num2
    if m > n {
        return findK(nums2, nums1, k)
    }

    // 基础情况
    if len(nums1) == 0 {
        return nums2[k-1]
    }

    // 退出条件
    if k == 1 {
        return min(nums1[0], nums2[0])
    }

    // 找出两个数组中第k/2个数字
    index1 := min(m, k/2) - 1
    index2 := min(n, k/2) - 1
    // 如果下面成立，证明nums2中对应的数字更小，那么中位数一定会出现在nums1里，那么就把num2的前面k/2排除掉
    if nums1[index1] > nums2[index2] {
        return findK(nums1, nums2[index2+1:], k-index2-1)
    } else {
        return findK(nums1[index1+1:], nums2, k-index1-1)
    }
}


func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}
