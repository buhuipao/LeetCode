// 两种方法：
// 1）宽度优先topSort，逐步弹出节点，直到发现有环，然后选取最后一个节点返回
// 2）并查集，沿着给出的边联通各个节点，如果发现已经联通的边再次出现，则证明是环中最后一个节点
func findRedundantConnection(edges [][]int) []int {
    // 初始化
    parent := make([]int, len(edges)+1)
    for i := range parent {
        parent[i] = i
    }
    // 定义并查操作
    var find func(int) int
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }
    union := func(from, to int) bool {
        x, y := find(from), find(to)
        // 如果两个节点原来就已经联通，则证明找到最终的节点了
        if x == y {
            return true
        }
        parent[x] = y
        return false    // 两个节点之前不联通
    }
    for _, e := range edges {
        // 之前就已经联通，则返回答案
        if union(e[0], e[1]) {
            return e
        }
    }
    return nil
}
