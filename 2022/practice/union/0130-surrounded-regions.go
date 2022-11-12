func solve(board [][]byte)  {
    rows := len(board)
    if rows == 0 {
        return
    }

    cols := len(board[0])

    u := NewUF(rows*cols+1)
    // 把所有边缘节点给末尾的虚拟节点连通
    dummy := rows*cols
    for i := 0; i < rows; i++ {
        if board[i][0] == 'O' {
            u.union(i*cols+0, dummy)
        }

        if board[i][cols-1] == 'O' {
            u.union(i*cols+cols-1, dummy)
        }
    }

    for j := 0; j < cols; j++ {
        if board[0][j] == 'O' {
            u.union(0*rows+j, dummy)
        }

        if board[rows-1][j] == 'O' {
            u.union((rows-1)*cols+j, dummy)
        }
    }

    d := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    for i := 0; i < rows-1; i++ {
        for j := 0; j < cols-1; j++ {
            if board[i][j] == 'O' {
                for _, dv := range d {
                    x, y := i+dv[0], j+dv[1]
                    // 判断边界
                    if x < 0 || x >= rows || y < 0 || y >= cols {
                        continue
                    }

                    // 将(i, j) 与上下左右连接
                    if board[x][y] == 'O' {
                        u.union(x*cols+y, i*cols+j)
                    }
                }
            }
        }
    }

    for i := 0; i < rows-1; i++ {
        for j := 0; j < cols-1; j++ {
            if !u.connected(i*cols+j, dummy) {
                board[i][j] = 'X'
            }
        }
    }
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
