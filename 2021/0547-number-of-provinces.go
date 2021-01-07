// 并查集
func findCircleNum(isConnected [][]int) (ans int) {
    n := len(isConnected)
    // 初始化
    parent := make([]int, n)
    for i := range parent {
        parent[i] = i
    }
    var find func(int) int
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }
    union := func(from, to int) {
        parent[find(from)] = find(to)
    }
    // 处理所有节点
    for i, row := range isConnected {
        for j := i + 1; j < n; j++ {
            if row[j] == 1 {
                union(i, j)
            }
        }
    }
    // 遍历父节点 
    for i, p := range parent {
        // 自身是root节点
        if i == p {
            ans++
        }
    }
    return
}
