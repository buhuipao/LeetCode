func findOrder(numCourses int, prerequisites [][]int) []int {
    // 建立出度表和入度计数
    outDegree, inDegreeCount := make(map[int][]int), make([]int, numCourses)
    for i := range prerequisites {
        des, src := prerequisites[i][0], prerequisites[i][1]
        outDegree[src] = append(outDegree[src], des)
        inDegreeCount[des] += 1
    }

    // 找到入度为0的节点
    stack := make([]int, 0, numCourses)
    for des, v := range inDegreeCount {
        if v == 0 {
            stack = append(stack, des)
        }
    }

    // 进行拓扑排序
    var ans []int
    var node int
    for len(stack) != 0 {
        stack, node = stack[:len(stack)-1], stack[len(stack)-1]
        ans = append(ans, node)

        for _, out := range outDegree[node] {
            inDegreeCount[out] -= 1
            if inDegreeCount[out] == 0 {
                stack = append(stack, out)
            }
        }
    }
    
    if len(ans) != numCourses {
        return nil
    }
    
    return ans
}
