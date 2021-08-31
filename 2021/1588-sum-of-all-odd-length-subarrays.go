// 按每个位置的贡献值计算
func sumOddLengthSubarrays(arr []int) int {
    n := len(arr)
    var ans int
    for i := 0; i < n; i++ {
        l1, r1 := (i + 1) / 2, (n - i) / 2 // 奇数
        l2, r2 := i / 2, (n - i - 1) / 2 // 偶数
        l2++
        r2++
        ans += (l1 * r1 + l2 * r2) * arr[i]
    }

    return ans;
}
