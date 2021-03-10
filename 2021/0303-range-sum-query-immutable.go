type NumArray struct {
    data []int
    sum []int
}


func Constructor(nums []int) NumArray {
    n := len(nums)
    obj := NumArray{data: nums, sum: make([]int, n)}
    if n == 0 {
        return obj
    }
    obj.sum[0] = nums[0]
    for i := 1; i < n; i++ {
        obj.sum[i] = obj.sum[i-1] + nums[i]
    }
    return obj
}


func (this *NumArray) SumRange(i int, j int) int {
    if i == 0 {
        return this.sum[j]
    }
    return this.sum[j] - this.sum[i-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
