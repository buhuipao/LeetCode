// 并查集：统计出整个石头有多少堆，每个堆必然会剩下一颗无法消除，所以答案就是：总数-堆数
func removeStones(stones [][]int) int {
    fa, rank := map[int]int{}, map[int]int{}
    var find func(int) int
    find = func(x int) int {
        // 初始化父亲和重量
        if _, ok := fa[x]; !ok {
            fa[x], rank[x] = x, 1
        }
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
        // 比较权重
        if rank[fx] < rank[fy] {
            fx, fy = fy, fx
        }
        rank[fx] += rank[fy]
        fa[fy] = fx
    }

    for _, p := range stones {
        // 这里把y增加10001是为了错开x与y，这样就可以把x、y轴一视同仁处理更简单
        union(p[0], p[1]+10001)
    }
    ans := len(stones)
    for x, fx := range fa {
        if x == fx {
            ans--
        }
    }
    return ans
}
