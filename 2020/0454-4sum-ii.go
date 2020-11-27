// 分组后就是两数之和
// 注意：可能出现多对
func fourSumCount(a, b, c, d []int) (ans int) {
    countAB := map[int]int{}
    for _, v := range a {
        for _, w := range b {
            countAB[v+w]++
        }
    }
    for _, v := range c {
        for _, w := range d {
            ans += countAB[-v-w]
        }
    }
    return
}
