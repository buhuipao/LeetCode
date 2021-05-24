func xorQueries(arr []int, queries [][]int) []int {
    n := len(arr)
    for i := 1; i < n; i++ {
        arr[i] ^= arr[i-1]
    }
    ans := make([]int, len(queries))
    for i, q := range queries {
        if q[0] == 0 {
            ans[i] = arr[q[1]]
        } else {
            ans[i] = arr[q[0]-1] ^ arr[q[1]]
        }
    }
    return ans
}
