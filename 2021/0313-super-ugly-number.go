func nthSuperUglyNumber(n int, primes []int) (ugly int) {
    seen := map[int]bool{1: true}
    h := &hp{[]int{1}}
    for i := 0; i < n; i++ {
        ugly = heap.Pop(h).(int)
        for _, prime := range primes {
            if next := ugly * prime; !seen[next] {
                seen[next] = true
                heap.Push(h, next)
            }
        }
    }
    return
}

// 堆
type hp struct{ sort.IntSlice }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
