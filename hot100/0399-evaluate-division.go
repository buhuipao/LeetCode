func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    u := UF{
        parent: make(map[string]string),
        value: make(map[string]float64),
        size: make(map[string]int),
        cnt: 0,
    }

    for i := range equations {
        a, b, val := equations[i][0], equations[i][1], values[i]
        // 以 b 为 parent
        u.union(a, b, val)
    }

    ans := make([]float64, len(queries))
    for i := range queries {
        _, ans[i] = u.connected(queries[i][0], queries[i][1])
    }

    return ans
}

type UF struct {
    parent map[string]string
    value map[string]float64
    size map[string]int
    cnt int
}

// find 找出父亲
func (u *UF) find(c string) string {
    base := 1.0
    root := c
    for u.parent[root] != root {
        root = u.parent[root]
        base *= u.value[root]
    }

    // 路径压缩
    // 同时更新自己相对于root的比例值
    node := c
    for node != root {
        nodeP := u.parent[node]
        // u.value[node]是指自己相对于parent的比例值，现在可能需要更新了
        u.value[node] *= base
        base /= u.value[nodeP]
        // 压缩
        u.parent[node] = root
        node = nodeP
    }

    return root
}

// union 联接
func (u *UF) union(a, b string, value float64) {
    // 默认父母为自己
    if _, ok := u.parent[a]; !ok {
        u.parent[a] = a
        u.value[a] = 1.0
        u.size[a] = 1
    }

    if _, ok := u.parent[b]; !ok {
        u.parent[b] = b
        u.value[b] = 1.0
        u.size[b] = 1
    }

    parentA, parentB := u.find(a), u.find(b)
    if parentA != parentB {
        u.cnt--
        if u.size[parentA] > u.size[parentB] {
            u.parent[parentB] = parentA
            u.size[parentA] += u.size[parentB]
            u.value[parentB] = 1/value * u.value[a] / u.value[b]
        } else {
            u.parent[parentA] = parentB
            u.size[parentB] += u.size[parentA]
            u.value[parentA] = value * u.value[b] / u.value[a]
        }
    }
}

// connected 判断联接
func (u *UF) connected(a, b string) (bool, float64) {
    if _, ok := u.value[a]; !ok {
        return false, -1.0
    }

    if _, ok := u.value[b]; !ok {
        return false, -1.0
    }

    if ok := u.find(a) == u.find(b); !ok {
        return false, -1.0
    }

    return true, u.value[a] / u.value[b]
}

// count 统计
func (u *UF) count() int {
    return u.cnt
}
