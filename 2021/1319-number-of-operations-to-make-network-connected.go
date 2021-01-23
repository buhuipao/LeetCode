// 计算并查集的堆数，调整网线使其只能有一堆
func makeConnected(n int, connections [][]int) int {
    if n-1 > len(connections) {
        return -1
    }
    // 初始化并查集
    parent := make([]int, n)
    for i := range parent {
        parent[i] = i 
    }
    var find func(int) int
    find = func(x int) int {
        if x != parent[x] {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }
    union := func(x, y int) bool {
        a, b := find(x), find(y)
        if a != b {
            parent[a] = b
            return true
        }
        return false
    }
    // 并查集操作
    for _, conn := range connections {
        union(conn[0], conn[1])
    }
    var count int
    for i := range parent {
        if i == parent[i] {
            count++
        }
    }
    return count-1
}
