// 找出边缘线出现的最大次数
func leastBricks(wall [][]int) int {
    n := len(wall)
    var max int
    m := make(map[int]int)
    for _, w := range wall {
        l := 0
        for _, num := range w[:len(w)-1] {
            l += num
            m[l] += 1
            if m[l] > max {
                max = m[l]
            } 
        }
    }
    return n - max
}
