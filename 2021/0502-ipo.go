//  避免死锁的银行家算法
func findMaximizedCapital(k, w int, profits, capital []int) int {
    n := len(profits)
    type pair struct{ c, p int }
    arr := make([]pair, n)
    for i, p := range profits {
        arr[i] = pair{capital[i], p}
    }
    sort.Slice(arr, func(i, j int) bool { return arr[i].c < arr[j].c })

    h := &hp{}
    for cur := 0; k > 0; k-- {
        for cur < n && arr[cur].c <= w {
            heap.Push(h, arr[cur].p)
            cur++
        }
        if h.Len() == 0 {
            break
        }
        w += heap.Pop(h).(int)
    }
    return w
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
