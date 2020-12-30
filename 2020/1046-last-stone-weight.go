type hp struct{ sort.IntSlice }
// 实现堆的三个接口
func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }   // 覆盖sort.IntSlice的Less接口；设成反向的，因为需要构建大根堆
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
// 自用两个接口
func (h *hp) push(v int)         { heap.Push(h, v) }
func (h *hp) pop() int           { return heap.Pop(h).(int) }

// 使用heap去实现，模拟整个过程即可
func lastStoneWeight(stones []int) int {
    q := &hp{stones}
    heap.Init(q)
    for q.Len() > 1 {
        x, y := q.pop(), q.pop()
        // x 必定大于等于 y
        if x > y {
            q.push(x - y)
        }
        // else 条件是等于所以不做处理
    }
    if q.Len() > 0 {
        return q.IntSlice[0]
    }
    return 0
}
