type Solution struct {
    pre []int
}

func Constructor(w []int) Solution {
    for i := 1; i < len(w); i++ {
        w[i] += w[i-1]
    }
    return Solution{w}
}

func (s *Solution) PickIndex() int {
    x := rand.Intn(s.pre[len(s.pre)-1]) + 1
    return sort.SearchInts(s.pre, x)
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
