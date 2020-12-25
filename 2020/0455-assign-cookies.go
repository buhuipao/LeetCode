// 尽可能满足小胃口的就好了的，大胃口的不管了
func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    n :=  len(s)
    var ans, j int
    for i := range g {
        for j < n {
            if g[i] <= s[j] {
                j++
                ans++
                break
            }
            j++
        }
        // 提前返回
        if j == n {
            return ans
        }
    }
    return ans
}
