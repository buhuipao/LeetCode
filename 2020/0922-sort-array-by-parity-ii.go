// 双指针
func sortArrayByParityII(a []int) []int {
    for i, j := 0, 1; i < len(a); i += 2 {
        // 如果a[i]为奇数则进行交换
        if a[i]%2 == 1 {
            //  找到a[j]为偶数的位置，并进行交换
            for a[j]%2 == 1 {
                j += 2
            }
            a[i], a[j] = a[j], a[i]
        }
    }
    return a
}

