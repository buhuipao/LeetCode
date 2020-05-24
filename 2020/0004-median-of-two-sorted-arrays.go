func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m, n := len(nums1), len(nums2)
    // 默认num1长度小于nums2长度
    if m > n {
        return findMedianSortedArrays(nums2, nums1)
    }
    // 中位数：(m + n + 1) / 2, (m + n + 2) / 2
    left, right := 0, m
    // value1为左边部分的最大值，value2为右边部分的最小值
    value1, value2 := 0, 0
    MinInt, MaxInt := ^(int(^uint(0)>>1)), int(^uint(0)>>1)
    var a, b, c, d int
    for left <= right {
        i := left + (right - left) / 2
        j := (m + n + 1) / 2 - i
        // 找出nums1[i-1], nums1[i], nums2[j-1], nums2[j]
        /* 为了简化分析，假设A[i−1],B[j−1],A[i],B[j] 总是存在。
        对于 i=0、i=m、j=0、j=n 这样的临界条件，我们只需要规定：A[−1]=B[−1]=−∞，A[m]=B[n]=∞ 即可。
        这也是比较直观的：
            1）当一个数组不出现在前一部分时，对应的值为负无穷，就不会对前一部分的最大值产生影响；
            2）当一个数组不出现在后一部分时，对应的值为正无穷，就不会对后一部分的最小值产生影响。
        */
        a = MinInt
        if i != 0 {
            a = nums1[i-1]
        }
        b = MaxInt
        if i != m {
            b = nums1[i]
        }
        c = MinInt
        if j != 0 {
            c = nums2[j-1]
        }
        d = MaxInt
        if j != n {
            d = nums2[j]
        }
        if a <= d {
            left = i + 1
            value1, value2 = max(a, c), min(b, d)
        } else {
            right = i - 1
        }
    }
    // 区分两个数组的总长度
    if (m + n) % 2 == 1 {
        return float64(value1)
    } else {
        return (float64(value1) + float64(value2)) / 2
    }
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
