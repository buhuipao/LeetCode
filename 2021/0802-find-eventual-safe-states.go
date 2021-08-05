// 反图 + 拓扑排序
func eventualSafeNodes(graph [][]int) (ans []int) {
    n := len(graph)
    rg := make([][]int, n)
    inDeg := make([]int, n)
    for x, ys := range graph {
        for _, y := range ys {
            rg[y] = append(rg[y], x)
        }
        inDeg[x] = len(ys)
    }

    q := []int{}
    for i, d := range inDeg {
        if d == 0 {
            q = append(q, i)
        }
    }
    for len(q) > 0 {
        y := q[0]
        q = q[1:]
        for _, x := range rg[y] {
            inDeg[x]--
            if inDeg[x] == 0 {
                q = append(q, x)
            }
        }
    }

    for i, d := range inDeg {
        if d == 0 {
            ans = append(ans, i)
        }
    }
    return
}
