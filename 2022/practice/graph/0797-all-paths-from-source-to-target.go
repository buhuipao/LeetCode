func allPathsSourceTarget(graph [][]int) [][]int {
    var ans [][]int
    traverse(graph, 0, []int{}, &ans)
    return ans
}

func traverse(graph [][]int, i int, path []int, ans *[][]int) {
    path = append(path, i)
    n := len(graph)
    if i == n-1 {
        tmp := make([]int, len(path))
        copy(tmp, path)
        *ans = append(*ans, tmp)
    }

    for _, j := range graph[i] {
        traverse(graph, j, path, ans)
    }
}
