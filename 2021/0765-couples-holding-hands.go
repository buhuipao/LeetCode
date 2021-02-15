type unionFind struct {
    parent, size []int
    setCount     int // 当前连通分量数目
}

func newUnionFind(n int) *unionFind {
    parent := make([]int, n)
    size := make([]int, n)
    for i := range parent {
        parent[i] = i
        size[i] = 1
    }
    return &unionFind{parent, size, n}
}

func (uf *unionFind) find(x int) int {
    if uf.parent[x] != x {
        uf.parent[x] = uf.find(uf.parent[x])
    }
    return uf.parent[x]
}

func (uf *unionFind) union(x, y int) {
    fx, fy := uf.find(x), uf.find(y)
    if fx == fy {
        return
    }
    if uf.size[fx] < uf.size[fy] {
        fx, fy = fy, fx
    }
    uf.size[fx] += uf.size[fy]
    uf.parent[fy] = fx
    uf.setCount--
}

// 可以把一对情侣看做是一个节点，把座位弄错的情侣(s)理解成相互连接，此时要做的工作是把所有连接剔除，
// 而每个N节点的联通图需要剔除的步骤是N-1，所以最终需要交换的次数为：N - 连通分量
func minSwapsCouples(row []int) int {
    n := len(row)
    uf := newUnionFind(n / 2)
    for i := 0; i < n; i += 2 {
        uf.union(row[i]/2, row[i+1]/2)
    }
    return n/2 - uf.setCount
}
