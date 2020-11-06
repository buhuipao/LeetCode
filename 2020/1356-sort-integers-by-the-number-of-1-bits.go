// 分两步：
// 1）计算每个数字的1个数，并放到桶里
// 2）分别排序每个桶里数字，并追加到ans里
func sortByBits(arr []int) []int {
    m := make([][]int, 32)
    for i := range m {
        m[i] = make([]int, 0)
    }
    for i := range arr {
        c := count(arr[i])
        m[c]  = append(m[c], arr[i])
    }
    ans := make([]int, 0, len(arr))
    for i := range m {
        sort.Ints(m[i])
        ans = append(ans, m[i]...)
    }
    return ans
}

func count(a int) int {
    var c int
    for a != 0 {
        if a % 2 == 1 {
            c++
        }
        a >>= 1
    }
    return c 
}
