// 并查集
func equationsPossible(equations []string) bool {
    u := NewUF(26)
    // 现将等式左右节点联接起来
    var x, y int
    for i := range equations {
        if equations[i][1:3] == "==" {
            x, y = int(equations[i][0]-'a'), int(equations[i][3]-'a')
            u.union(x, y)
        }
    }

    // 判定不等式左右节点是否已联接
    for i := range equations {
        if equations[i][1:3] == "!=" {
            x, y = int(equations[i][0]-'a'), int(equations[i][3]-'a')
            if u.connected(x, y) {
                return false
            }
        }
    }

    return true
}


type UF struct {
    cnt int
    parent []int
    weight []int
}

func NewUF(n int) UF {
    weight := make([]int, n)
    for i := range weight {
        weight[i] = 1
    }

    parent := make([]int, n)
    for i := range parent {
        parent[i] = i
    }

   return UF{cnt: n, weight: weight, parent: parent} 
}

func (u UF) union(p, q int) {
    rootQ := u.find(q)
    rootP := u.find(p)
    if rootP == rootQ {
        return 
    }

    // 把小树接到大树上
    if u.weight[rootP] > u.weight[rootQ] {
        u.parent[rootQ] = rootP
        u.weight[rootP] += u.weight[rootQ]
    } else {
        u.parent[rootP] = rootQ
        u.weight[rootQ] += u.weight[rootP]
    }

    u.cnt--
}

func (u UF) connected(p, q int) bool {
    return u.find(p) == u.find(q)
}

func (u UF) find(p int) int {
    for p != u.parent[p] {
        // 路径压缩
        u.parent[p] = u.parent[u.parent[p]]
        p = u.parent[p]
    }

    return u.parent[p]
}

func (u UF) count() int {
    return u.cnt
}
