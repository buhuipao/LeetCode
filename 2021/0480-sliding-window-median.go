// 大小根堆
type hp struct {
    sort.IntSlice
    size int
}
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp) push(v int)         { h.size++; heap.Push(h, v) }
func (h *hp) pop() int           { h.size--; return heap.Pop(h).(int) }
func (h *hp) prune() {
    for h.Len() > 0 {
        num := h.IntSlice[0]
        if h == small {
            num = -num
        }
        if d, has := delayed[num]; has {
            if d > 1 {
                delayed[num]--
            } else {
                delete(delayed, num)
            }
            heap.Pop(h)
        } else {
            break
        }
    }
}

var delayed map[int]int
var small, large *hp

func medianSlidingWindow(nums []int, k int) []float64 {
    delayed = map[int]int{} // 哈希表，记录「延迟删除」的元素，key 为元素，value 为需要删除的次数
    small = &hp{}           // 大根堆，维护较小的一半元素
    large = &hp{}           // 小根堆，维护较大的一半元素
    makeBalance := func() {
        // 调整 small 和 large 中的元素个数，使得二者的元素个数满足要求
        if small.size > large.size+1 { // small 比 large 元素多 2 个
            large.push(-small.pop())
            small.prune() // small 堆顶元素被移除，需要进行 prune
        } else if small.size < large.size { // large 比 small 元素多 1 个
            small.push(-large.pop())
            large.prune() // large 堆顶元素被移除，需要进行 prune
        }
    }
    insert := func(num int) {
        if small.Len() == 0 || num <= -small.IntSlice[0] {
            small.push(-num)
        } else {
            large.push(num)
        }
        makeBalance()
    }
    erase := func(num int) {
        delayed[num]++
        if num <= -small.IntSlice[0] {
            small.size--
            if num == -small.IntSlice[0] {
                small.prune()
            }
        } else {
            large.size--
            if num == large.IntSlice[0] {
                large.prune()
            }
        }
        makeBalance()
    }
    getMedian := func() float64 {
        if k&1 > 0 {
            return float64(-small.IntSlice[0])
        }
        return float64(-small.IntSlice[0]+large.IntSlice[0]) / 2
    }

    for _, num := range nums[:k] {
        insert(num)
    }
    n := len(nums)
    ans := make([]float64, 1, n-k+1)
    ans[0] = getMedian()
    for i := k; i < n; i++ {
        insert(nums[i])
        erase(nums[i-k])
        ans = append(ans, getMedian())
    }
    return ans
}
