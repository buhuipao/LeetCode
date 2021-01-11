// 并查集
func smallestStringWithSwaps(s string, pairs [][]int) string {
    n := len(s)
    fa := make([]int, n)
    rank := make([]int, n)
    for i := range fa {
        fa[i] = i
        rank[i] = 1
    }
    var find func(int) int
    find = func(x int) int {
        if fa[x] != x {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }
    union := func(x, y int) {
        fx, fy := find(x), find(y)
        if fx == fy {
            return
        }
        if rank[fx] < rank[fy] {
            fx, fy = fy, fx
        }
        rank[fx] += rank[fy]
        fa[fy] = fx
    }

    for _, p := range pairs {
        union(p[0], p[1])
    }

    groups := map[int][]byte{}
    for i := range s {
        f := find(i)
        groups[f] = append(groups[f], s[i])
    }

    for _, bytes := range groups {
        sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
    }

    ans := make([]byte, n)
    for i := range ans {
        f := find(i)
        ans[i] = groups[f][0]
        groups[f] = groups[f][1:]
    }
    return string(ans)
}
