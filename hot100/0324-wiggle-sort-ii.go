func wiggleSort(nums []int)  {
    arr := append([]int{}, nums...)
    sort.Ints(arr)

    n := len(nums)
    x := (n + 1) / 2
    // nums[0], nums[1], ..., nums[x-1], nums[x], ... nums[n]
    // 为了避免将x-1和x混在一起导致错误，于是需要将它们两尽可能分开，最好是一个在队首一个在队尾
    // 一个行之有效的办法是：将前后部分反转后再merge，保持nums[x]在队首，nums[x-1]在队尾
    for i, j, k := 0, x-1, n-1; i < n; i += 2 {
        nums[i] = arr[j]
        if i + 1 < n {
            nums[i+1] = arr[k]
        }

        j--
        k--
    }
}
