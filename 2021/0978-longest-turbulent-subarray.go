func maxTurbulenceSize(arr []int) int {
    n := len(arr)
    var ans, tmp int
    for i := 0; i < n; i++ {
        var v1, v2 int
        // 相邻的三个数的斜率相反
        if i > 0 {
            v1 = arr[i] - arr[i-1]
        }
        if i < n-1 {
            v2 = arr[i+1] - arr[i]
        }
        if v1 * v2 < 0 {    // 连续的三个数形成山顶或者山谷
            tmp++
        } else if v1 != 0 || v2 != 0 {  // 递增或者递减，形成了山坡
            tmp = 2
        } else {    // 只有一个数或者所有的数相同，也就是说只形成了一个点或者平台
            tmp = 1
        }
        ans = max(tmp, ans)
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
