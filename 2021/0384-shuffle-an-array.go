import "math/rand"

type Solution struct {
    raw []int
    cur []int
}


func Constructor(nums []int) Solution {
    // 加上随机种子
    rand.Seed(time.Now().UnixNano())
    cur := make([]int, len(nums))
    copy(cur, nums)
    return Solution{raw: nums, cur: cur}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
    copy(this.cur, this.raw)
    return this.raw
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
    for i := len(this.raw) - 1; i >= 0; i-- {
        // 没有seed那一行，这一行将不生效
        index := rand.Intn(i+1)
        // 随机选择一个index与当前节点互d换
        this.cur[index], this.cur[i] = this.cur[i], this.cur[index]
    }
    return this.cur
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
