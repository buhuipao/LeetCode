type CQueue struct {
    s1 []int
    s2 []int
}


func Constructor() CQueue {
    return CQueue{s1: []int{}, s2: []int{}}
}


func (this *CQueue) AppendTail(value int)  {
    this.s1 = append(this.s1, value)
    this.adjust()
}


func (this *CQueue) DeleteHead() int {
    this.adjust()
    value := -1
    if len(this.s2) > 0 {
        value = this.s2[len(this.s2)-1]
        this.s2 = this.s2[:len(this.s2)-1]
    }
    return value
}

// 调整两个栈，确保第二个栈（输出栈）不为空
func (this *CQueue) adjust() {
    if len(this.s2) == 0 {
        for i := len(this.s1)-1; i >= 0; i-- {
            this.s2 = append(this.s2, this.s1[i])
        }
        this.s1 = this.s1[:0]
    }
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
