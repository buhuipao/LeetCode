type MaxQueue struct {
    data []int
    max []int
}


func Constructor() MaxQueue {
    return MaxQueue{data: make([]int, 0), max: make([]int, 0)}
}


func (this *MaxQueue) Max_value() int {
    if len(this.max) == 0 {
        return -1
    }
    return this.max[0]
}


func (this *MaxQueue) Push_back(value int)  {
    this.data = append(this.data, value)
    for len(this.max) > 0 && this.max[len(this.max)-1] < value {
        this.max = this.max[:len(this.max)-1]
    }
    this.max = append(this.max, value)
}


func (this *MaxQueue) Pop_front() int {
    if len(this.data) == 0 {
        return -1
    }
    value := this.data[0]
    this.data = this.data[1:]
    if value == this.max[0] {
        this.max = this.max[1:]
    }
    return value
}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
