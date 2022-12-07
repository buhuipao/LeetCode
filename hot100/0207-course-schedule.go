func canFinish(numCourses int, prerequisites [][]int) bool {
    // 建立出度和入度统计
    outDegree := make(map[int][]int, numCourses)
    inDegree := make([]int, numCourses)
    for i := range prerequisites {
        des, src := prerequisites[i][0], prerequisites[i][1]
        outDegree[src] = append(outDegree[src], des)
        inDegree[des] += 1
    }

    // 收集入度为0的节点
    var stack []int
    for i := range inDegree {
        if inDegree[i] == 0 {
            stack = append(stack, i)
        }
    }

    // 拓扑排序，收集可遍历的节点
    var n, node int
    for len(stack) != 0 {
        node, stack = stack[len(stack)-1], stack[:len(stack)-1]
        n++

        for _, des := range outDegree[node] {
            inDegree[des] -= 1
            if inDegree[des] == 0 {
                stack = append(stack, des)
            }
        }
    }

    return n == numCourses
}
