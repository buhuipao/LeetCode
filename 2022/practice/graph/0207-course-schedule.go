// 拓扑排序
func canFinish(numCourses int, prerequisites [][]int) bool {
    // 建立出度表以及入度统计
    outDegree := make(map[int][]int, numCourses)
    count := make([]int, numCourses)
    for i := range prerequisites {
        src, des := prerequisites[i][1], prerequisites[i][0]
        outDegree[src] = append(outDegree[src], des)
        count[des] += 1
    }

    // 收集入度为0的节点
    var stack []int
    for des := range count {
        if count[des] == 0 {
            stack = append(stack, des)
        }
    }

    // 拓扑排序，统计可遍历的节点
    var n, node int
    for len(stack) != 0 {
        node, stack = stack[len(stack)-1], stack[:len(stack)-1]
        n++
        
        for _, v := range outDegree[node] {
            count[v] -= 1
            if count[v] == 0 {
                stack = append(stack, v)
            }
        }
    }

    return numCourses == n
}
