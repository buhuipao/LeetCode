type KthLargest struct {
    sort.IntSlice
    k int
}


func Constructor(k int, nums []int) KthLargest {
    o := KthLargest{k: k}
    for _, v := range nums {
        o.Add(v)
    }

    return o
}

func (o *KthLargest) Push(v interface{}) {
    o.IntSlice = append(o.IntSlice, v.(int))
}

func (o *KthLargest) Pop() interface{} {
    v := o.IntSlice[len(o.IntSlice)-1]
    o.IntSlice = o.IntSlice[:len(o.IntSlice)-1]
    return v
}


func (this *KthLargest) Add(val int) int {
    heap.Push(this, val)
    if len(this.IntSlice) > this.k {
        heap.Pop(this)
    }
    
    return this.IntSlice[0]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
