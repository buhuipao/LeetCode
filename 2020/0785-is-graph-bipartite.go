var (
    RED, GREEN, UNCOLORED int = -1, 1, 0
    color []int
    valid bool
)


func isBipartite(graph [][]int) bool {
    n := len(graph)
    color = make([]int, n)
    valid = true
    for i := 0; i < n && valid; i++ {
        if color[i] == UNCOLORED {
            dfs(i, RED, graph)
        }
    }
    return valid
}

func dfs(node, c int, graph [][]int) {
    color[node] = c
    // 相反的颜色
    C := c * -1
    for _, neighbor := range graph[node] {
        if color[neighbor] == UNCOLORED {
            dfs(neighbor, C, graph)
        // 如果有颜色，还与预期颜色不符合，那么直接设置为非法的
        } else if color[neighbor] != C {
            valid = false
            return
        }
    }
}
