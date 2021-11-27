type Solution struct {
    m, n, total int
    mp          map[int]int
}

func Constructor(m, n int) Solution {
    return Solution{m, n, m * n, map[int]int{}}
}

func (s *Solution) Flip() (ans []int) {
    x := rand.Intn(s.total)
    s.total--
    if y, ok := s.mp[x]; ok { // 查找位置 x 对应的映射
        ans = []int{y / s.n, y % s.n}
    } else {
        ans = []int{x / s.n, x % s.n}
    }
    if y, ok := s.mp[s.total]; ok { // 将位置 x 对应的映射设置为位置 total 对应的映射
        s.mp[x] = y
    } else {
        s.mp[x] = s.total
    }
    return
}

func (s *Solution) Reset() {
    s.total = s.m * s.n
    s.mp = map[int]int{}
}
