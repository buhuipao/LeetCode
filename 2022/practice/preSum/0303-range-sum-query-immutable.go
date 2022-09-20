type NumArray struct {
    data []int
}


func Constructor(nums []int) NumArray {
    o := NumArray{
        data: make([]int, len(nums)+1),
    }
    v := o.data[0]
    for i := range nums {
        v += nums[i]
        o.data[i+1] = v 
    }

    return o
}


func (this *NumArray) SumRange(left int, right int) int {
    return this.data[right+1] - this.data[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
