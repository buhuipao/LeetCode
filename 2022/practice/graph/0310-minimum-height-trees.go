// 两种方法：
// 1）找到图中的最长路径，处于路径上中间位置的节点可以使得高度为路径的二分之一
// 2）拓扑排序
func findMinHeightTrees(n int, edges [][]int) []int {
    if n == 1 {
        return []int{0}
    }

    // 建立邻接表以及度统计
    g := make([][]int, n)
    deg := make([]int, n)
    for i := range edges {
        src, des := edges[i][0], edges[i][1]
        g[src] = append(g[src], des)
        g[des] = append(g[des], src)
        deg[src]++
        deg[des]++
    }

    // 统计度为1的节点，这种叶子结点一定是在最长路径两端
    var q []int
    for i := range deg {
        if deg[i] == 1 {
            q = append(q, i)
        }
    }

    // 进行拓扑排序
    remainNodeCnt := n 
    for remainNodeCnt > 2 {
        remainNodeCnt -= len(q)
        tmp := q
        q = nil

        for _, v := range tmp {
            for _, n := range g[v] {
                deg[n]--
                if deg[n] == 1 {
                    q = append(q, n)
                }
            }
        }
    }

    return q
}
